package user

import (
	"database/sql"
	"fmt"
	"login_page_gerin/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func InitUserRoute(mainRoute string, db *sql.DB, r *mux.Router) {
	UserController := NewController(db)
	p := r.PathPrefix(mainRoute).Subrouter()
	p.Use(middleware.TokenValidation)
	r.HandleFunc("/user/login", UserController.HandleUserLogin()).Methods("POST")
	r.HandleFunc("/user/register", UserController.HandleRegisterNewUser()).Methods("POST")
	r.HandleFunc("/register", registerPage).Methods("GET")
	p.HandleFunc("/list", userList).Methods("GET")
	p.HandleFunc("", UserController.HandleGETAllUsers()).Methods("GET")
	p.HandleFunc("/logout", UserController.UserLogOut()).Methods("GET")
	p.HandleFunc("/{id}", UserController.HandleUPDATEUsers()).Methods("PUT")
	p.HandleFunc("/{id}", UserController.HandleDELETEUsers()).Methods("DELETE")
}

func registerPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<!DOCTYPE html>
	<html lang="en">
	  <head>
		<meta charset="UTF-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>Document</title>
	  </head>
	  <body>
		<h1>Register Page</h1>
		<form action="http://localhost:8080/user/register" method="POST">
		  <input class="input100" type="text" name="username" placeholder="Username" />
		  <input class="input100" type="text" name="password" placeholder="Password" />
		  <input class="input100" type="text" name="alamat" placeholder="Alamat" />
		  <input class="input100" type="text" name="telp" placeholder="Telp" />
		  <button class="login100-form-btn">
			Register
		  </button>
		</form>
	  </body>
	</html>
	`)
}

func userList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Document</title>
    <style>
      table,
      th,
      td {
        border: 1px solid black;
      }
      table {
        border-collapse: collapse;
        width: 60%;
      }

      th {
        height: 50px;
      }
      th,
      td {
        padding: 15px;
        text-align: left;
        border-bottom: 1px solid #ddd;
      }
      tr:hover {
        background-color: #f5f5f5;
      }
    </style>
  </head>
  <body>
    <h1>Daftar User</h1>
	<button class="btn btn-success" onclick="location.href='http://localhost:8080/user/logout';"> Logout</button>
    <table id="myTable">
      <tr>
        <td><b>Username</b></td>
        <td><b>Alamat</b></td>
        <td><b>Telp</b></td>
      </tr>
    </table>
    <script>
      async function getUserAsync(name) {
        try {
          let response = await fetch("http://localhost:8080/user");
          let data = await response.json();
          getdata = data;
          return data;
        } catch (err) {
          console.log(err);
        }
      }

      let getdata = [];

      getUserAsync();
      setTimeout(() => {
        var table = document.getElementById("myTable");
        for (var i = 0; i < getdata.Results.length; ++i) {
          var row = table.insertRow(1);
          var cell1 = row.insertCell(0);
          var cell2 = row.insertCell(1);
          var cell3 = row.insertCell(2);

          cell1.innerHTML = getdata.Results[i].username;
          cell2.innerHTML = getdata.Results[i].alamat;
          cell3.innerHTML = getdata.Results[i].telp;
        }
	  }, 100);
	  document.addEventListener('DOMContentLoaded', getUserAsync, true);
    </script>
  </body>
</html>
`)
}
