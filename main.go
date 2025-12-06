package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/tdanieljr/gator/internal/config"
	"github.com/tdanieljr/gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}
type command struct {
	name string
	args []string
}
type commands struct {
	actions map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.actions[cmd.name]
	if !ok {
		return fmt.Errorf("Command not found: %s", cmd.name)
	}
	err := f(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
func (c *commands) register(name string, f func(*state, command) error) {
	c.actions[name] = f
}
func newCommands() commands {
	m := make(map[string]func(*state, command) error)
	cmds := commands{actions: m}
	return cmds
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("postgres", cfg.DBUrl)
	dbQueries := database.New(db)
	s := state{cfg: &cfg, db: dbQueries}
	cmds := newCommands()
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handdlerGetFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", handlerBrowse)
	input := os.Args
	if len(input) < 2 {
		fmt.Println(fmt.Errorf("Error: %s", err))

	}
	cmd := command{name: input[1], args: input[2:]}
	err = cmds.run(&s, cmd)
	if err != nil {
		panic(err)
	}

}
