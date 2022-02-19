package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/wonesy/bookfahrt/ent"
	"github.com/wonesy/bookfahrt/ent/club"
)

func (e *ApiEnv) InitClubRouter() func(fiber.Router) {
	return func(router fiber.Router) {
		router.Post("", e.CreateClubHandler())
		router.Get("/:id?", e.GetClubHandler())
		router.Put("/:id", e.UpdateClubHandler())
		router.Delete("/:id", e.DeleteClubHandler())
		router.Post("/:id/join", e.JoinClubHandler())
	}
}

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

func (e *ApiEnv) DeleteClub(id string) error {
	val, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return e.Client.Club.DeleteOneID(val).Exec(context.TODO())
}

func (e *ApiEnv) CreateClubHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var Club *ent.Club
		if err := c.BodyParser(Club); err != nil {
			return nil
		}

		user, err := e.GetSessionUser(c)
		if err != nil {
			return errors.Wrap(err, "CreateClubHandler")
		}

		createdClub, err := e.CreateClub(Club)
		if err != nil {
			return err
		}

		_, err = user.Update().
			AddClubs(createdClub).
			Save(context.Background())
		if err != nil {
			return errors.Wrap(err, "CreateClubHandler failed to add user to club")
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

		club, err := e.GetClubByName(name)
		if err != nil {
			return err
		}
		return c.JSON(club)
	}
}

func (e *ApiEnv) DeleteClubHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := e.DeleteClub(id)
		if err != nil {
			return err
		}
		return c.SendString("Deleted club")
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

func (e *ApiEnv) JoinClubHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		user, err := e.GetSessionUser(c)
		if err != nil {
			return errors.Wrap(err, "JoinClubHandler")
		}

		parsedID, err := uuid.Parse(id)
		if err != nil {
			return errors.Wrap(err, "JoinClubHandler parsing uuid failed")
		}

		_, err = e.UpdateUser(user, parsedID)
		return err
	}
}
