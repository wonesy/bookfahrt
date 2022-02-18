package api

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wonesy/bookfahrt/ent"
	"github.com/wonesy/bookfahrt/ent/club"
)

func (e *ApiEnv) GetAllClubs() ([]*ent.Club, error) {
	return e.Client.Club.Query().All(context.Background())
}

func (e *ApiEnv) GetClubByName(name string) (*ent.Club, error) {
	return e.Client.Club.Query().
		Where(club.NameEQ(name)).
		Only(context.Background())
}

func (e *ApiEnv) CreateClub(c *ent.Club) (*ent.Club, error) {
	return e.Client.Club.Create().
		SetName(c.Name).
		Save(context.Background())
}

func (e *ApiEnv) UpdateClub(c *ent.Club) (*ent.Club, error) {
	return nil, nil
}

func (e *ApiEnv) DeleteClub(name string) (int, error) {
	return 0, nil
}

func (e *ApiEnv) CreateClubHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var Club *ent.Club
		if err := c.BodyParser(Club); err != nil {
			return nil
		}
		createdClub, err := e.CreateClub(Club)
		if err != nil {
			return err
		}
		return c.JSON(createdClub)
	}
}

func (e *ApiEnv) GetClubHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		name := c.Params("name")
		if name == "" {
			Clubs, err := e.GetAllClubs()
			if err != nil {
				return err
			}
			return c.JSON(Clubs)
		}

		Club, err := e.GetClubByName(name)
		if err != nil {
			return err
		}
		return c.JSON(Club)
	}
}

func (e *ApiEnv) DeleteClubHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		name := c.Params("name")
		numDeleted, err := e.DeleteClub(name)
		if err != nil {
			return err
		}
		return c.SendString(fmt.Sprintf("Deleted %d club(s)", numDeleted))
	}
}

func (e *ApiEnv) UpdateClubHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var Club *ent.Club
		if err := c.BodyParser(Club); err != nil {
			return err
		}

		updatedClub, err := e.UpdateClub(Club)
		if err != nil {
			return err
		}

		return c.JSON(updatedClub)
	}
}
