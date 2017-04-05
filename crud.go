package main
//@author Lucas Martins Ramos

import (
	"log"

	//Para RESTful API
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/binding"

	//Para acesso ao database
	"database/sql"
	//Driver Mysql
	_ "github.com/go-sql-driver/mysql"
)

//Define objeto Product
type Product struct {
	Id_product int `json:"id_product"`
	Name string `json:"name" binding:"required"`
	Price float32 `json:"price"`
	Expiration_date  string `json:"expiration_date"`
	Owner string `json:"owner"`
}

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())

	log.Println("Tentando se concectar na base")
	//Configure aqui seu usuário e senha do Mysql (Deve ter permissão para crear shchema e tables,
	//para inicar a base aconselho usar root )
	db, err := sql.Open("mysql", "root:zxvf.321@tcp(localhost:3306)/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Criando database crud caso não exista")
	_,err = db.Exec("CREATE DATABASE IF NOT EXISTS crud")
    if err != nil {
       panic(err)
    }

    log.Println("Usando schema crud")
    _,err = db.Exec("USE crud")
    if err != nil {
       panic(err)
    }

    log.Println("Criando tabela Product caso não exita")
	_,err = db.Exec(`CREATE TABLE IF NOT EXISTS crud.product (
						id_product INT NOT NULL AUTO_INCREMENT,
				  		name VARCHAR(200) NULL,
				  		price DECIMAL(10,2) NULL,
				  		expiration_date DATE NULL,
							owner VARCHAR(20) NULL,
				  		PRIMARY KEY (id_product));`)
	if err != nil {
		log.Fatal(err)
	}



	//Controller para adicionar produto
	m.Post("/product/add/data.json", binding.Json(Product{}), func(product Product, r render.Render) {
		if err != nil {
			log.Println(err.Error())
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}
		stmt, err := db.Prepare(`Insert Into product
					 (name, price, expiration_date, owner)
					 Values (?, ?, ?, ?)`)
		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(product.Name, product.Price, product.Expiration_date, product.Owner)

		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}

		r.JSON(200, map[string]interface{}{"status": "success"})
	})

	//Controler para listar todos os produtos
	m.Get("/product/list/:owner", func(params martini.Params, r render.Render) {
		rows, err := db.Query(`Select id_product, name, price, expiration_date, owner
				       From product
							 WHERE owner = ?`, params["owner"])
		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}
		defer rows.Close()

		res := []Product{}
		for rows.Next() {
			p := Product{}
			rows.Scan(&p.Id_product, &p.Name, &p.Price, &p.Expiration_date, &p.Owner)
			res = append(res, p)
		}

		r.JSON(200, map[string]interface{}{"status": "success", "results": res})
	})

	//Controller para procurar por ID
	m.Get("/product/search/:id", func(params martini.Params, r render.Render) {

		stmt, err := db.Prepare(`Select id_product, name, price, expiration_date
					 From product
					 Where id_product = ?`)
		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}
		defer stmt.Close()

		p := Product{}
		err = stmt.QueryRow(params["id"]).Scan(&p.Id_product, &p.Name, &p.Price, &p.Expiration_date)

		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}

		r.JSON(200, map[string]interface{}{"status": "success", "results": p})
	})

	//Controller para buscar por Nome do produto
	m.Get("/product/searchByName/:name/:owner", func(params martini.Params, r render.Render) {

		stmt, err := db.Prepare(`Select id_product, name, price, expiration_date, owner
					 From product
					 Where name like ? AND owner = ?`)
		if err != nil {
			log.Println(err.Error())
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}
		defer stmt.Close()

		rows, err := stmt.Query(params["name"], params["owner"])

		res := []Product{}

		if err != nil {
			log.Println(err.Error())
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}

		for rows.Next() {
			p := Product{}
			rows.Scan(&p.Id_product, &p.Name, &p.Price, &p.Expiration_date, &p.Owner)
			res = append(res, p)
		}

		r.JSON(200, map[string]interface{}{"status": "success", "results": res})
	})

	//Controller para editar produto
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

	//Controller para excluir produto
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

	//Mapeamento de arquivos estáticos(HTML's, JS, CSS...)
	m.Use(martini.Static("static/"))

	m.Run()

}
