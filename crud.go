package main

import (
	//"fmt"
	"log"
	//for RESTful API
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render" //for rendering returning JSON
	//for sqlite3 storage, from drivers of go-wiki: https://code.google.com/p/go-wiki/wiki/SQLDrivers
	"database/sql"
	//use the initialization inside without actual use it: http://golang.org/doc/effective_go.html#blank
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Name string `json:"name" binding:"required"`
	Price float32 `json:"price"`
	Expiration_date  string `json:"expiration_date"`
}

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())

	log.Println("Tentando se concectar na base")
	//get database
	db, err := sql.Open("mysql", "root:lucasbemol@tcp(localhost:3306)/crud")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//using martini-contrib/binding for receiving incoming JSON
	m.Post("/product/add/data.json", binding.Json(Product{}), func(product Product, r render.Render) {
		log.Println("Valores no produto")
		log.Println(product.Name)
		log.Println(product.Price)
		log.Println(product.Expiration_date)

		stmt, err := db.Prepare(`Insert Into product
					 (name, price, expiration_date)
					 Values (?, ?, ?)`)
		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(product.Name, product.Price, product.Expiration_date)

		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}

		r.JSON(200, map[string]interface{}{"status": "success"})
	})

	m.Get("/product/list", func(r render.Render) {
		rows, err := db.Query(`Select name, price, expiration_date
				       From product`)
		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}
		defer rows.Close()

		//create a Struct array
		res := []Product{}
		for rows.Next() {
			//var tempID string
			p := Product{}
			rows.Scan(&p.Name, &p.Price, &p.Expiration_date)
			//append things after the array
			res = append(res, p)
		}

		r.JSON(200, map[string]interface{}{"status": "success", "results": res})
	})

	m.Get("/product/:id", func(params martini.Params, r render.Render) {
		//define SQL
		stmt, err := db.Prepare(`Select name, price, expiration_date
					 From product
					 Where id_product = ?`)
		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}
		defer stmt.Close()
		//execute with parameter
		p := Product{}
		err = stmt.QueryRow(params["id"]).Scan(&p.Name, &p.Price, &p.Expiration_date)

		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}

		r.JSON(200, map[string]interface{}{"status": "success", "results": p})
	})

	m.Put("/product/:id/data.json", binding.Json(Product{}), func(params martini.Params, product Product, r render.Render) {
		stmt, err := db.Prepare(`Update product
					 Set name=?, price=?, expiration_date=?
					 Where id_product=?`)
		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(product.Name, product.Price, product.Expiration_date, params["id"])

		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}
		r.JSON(200, map[string]interface{}{"status": "success"})
	})

	m.Delete("/product/delete/:id", func(params martini.Params, r render.Render) {
		stmt, err := db.Prepare(`Delete From product
					 Where id_product = ?`)
		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}
		defer stmt.Close()
		_, err = stmt.Exec(params["id"])
		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}

		r.JSON(200, map[string]interface{}{"status": "success"})
	})

	m.Run()
}