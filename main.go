package main

func main() {
	lattice := NewBlocklattice()
	defer lattice.db.Close()

	cli := CLI{lattice}
	cli.Run()
}