package api

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/wonesy/bookfahrt/ent"
	"github.com/wonesy/bookfahrt/ent/invitation"
)

func (e *ApiEnv) GetInvitationByID(id uuid.UUID) (*ent.Invitation, error) {
	return e.Client.Invitation.Query().
		Where(invitation.IDEQ(id)).
		WithClub().
		WithSponsor().
		Only(context.Background())
}

func (e *ApiEnv) CreateInvitation(sponsor *ent.User, clubID uuid.UUID) (*ent.Invitation, error) {
	return e.Client.Invitation.Create().
		SetClubID(clubID).
		SetSponsor(sponsor).
		Save(context.Background())
}

func (e *ApiEnv) UseAndDeleteInvitation(newUser *ent.User, invID uuid.UUID) error {
	ctx := context.Background()
	tx, err := e.Client.Tx(ctx)
	if err != nil {
		return errors.Wrap(err, "UseAndDeleteInvitation failed to create tx")
	}

	inv, err := e.GetInvitationByID(invID)
	if err != nil {
		return errors.Wrap(err, "UseAndDeleteInvitation no such invitation")
	}

	_, err = tx.User.UpdateOne(newUser).AddClubs(
		inv.Edges.Club,
	).Save(ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return err
	}

	if err := tx.Invitation.DeleteOneID(invID).Exec(ctx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return err
	}

	return tx.Commit()
}

func (e *ApiEnv) CreateInvitationHandler() func(*fiber.Ctx) error {
	type payload struct {
		ClubID string `json:"club_id"`
	}

	return func(c *fiber.Ctx) error {
		p := new(payload)
		if err := c.BodyParser(p); err != nil {
			return errors.Wrap(err, "CreateInvitationHandler failed to parse body")
		}

		user, err := e.GetSessionUser(c)
		if err != nil {
			return errors.Wrap(err, "CreateInvitationHandler")
		}

		clubUUID, err := uuid.Parse(p.ClubID)
		if err != nil {
			return errors.Wrap(err, "CreateInvitationHandler failed to parse club uuid")
		}

		inv, err := e.CreateInvitation(user, clubUUID)
		if err != nil {
			return errors.Wrap(err, "CreateInvitationHandler failed to create invitation")
		}

		return c.JSON(inv)
	}
}
