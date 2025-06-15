package main

import (
	"be-inventaris/barang"
	"be-inventaris/kategori"
	"be-inventaris/merk"
	"be-inventaris/peminjaman"
	"be-inventaris/ruangan"
	"be-inventaris/service"
	"be-inventaris/status"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type User struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Status   string `json:"status"`
	Password string `json:"password,omitempty"` // disembunyikan saat dikirim ke frontend
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/auth/login", LoginHandler).Methods("POST")
	router.HandleFunc("/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users", CreateUser).Methods("POST")

	//ROUTE: kategori barang (package kategori barang)
	router.HandleFunc("/kategori", kategori.GetKategori).Methods("GET")
	router.HandleFunc("/kategori", kategori.TambahKategori).Methods("POST")
	router.HandleFunc("/kategori/{id}", kategori.HapusKategori).Methods("DELETE")

	// routing merk
	router.HandleFunc("/merk", merk.GetMerk).Methods("GET")
	router.HandleFunc("/merk", merk.TambahMerk).Methods("POST")
	router.HandleFunc("/merk/{id}", merk.HapusMerk).Methods("DELETE")

	// router baru untuk status barang
	router.HandleFunc("/status", status.GetStatusSummary).Methods("GET")
	router.HandleFunc("/status/filter", status.GetBarangByStatus).Methods("GET")

	// router baru untuk peminjaman
	router.HandleFunc("/peminjaman", peminjaman.GetAllPeminjaman).Methods("GET")
	router.HandleFunc("/peminjaman", peminjaman.TambahPeminjaman).Methods("POST")
	router.HandleFunc("/peminjaman", peminjaman.EditPeminjaman).Methods("PUT")
	router.HandleFunc("/peminjaman", peminjaman.HapusPeminjaman).Methods("DELETE")

	// router untuk service barang
	router.HandleFunc("/service", service.GetServiceBarang).Methods("GET")
	router.HandleFunc("/service", service.TambahServiceBarang).Methods("POST")
	router.HandleFunc("/service", service.HapusServiceBarang).Methods("DELETE")
	router.HandleFunc("/service/selesai", service.SelesaiServiceBarang).Methods("PUT")

	// router untuk ruangan
	router.HandleFunc("/ruangan", ruangan.GetRuangan)             // GET
	router.HandleFunc("/ruangan/tambah", ruangan.TambahRuangan)   // POST
	router.HandleFunc("/ruangan/barang", ruangan.BarangByRuangan) // GET ?id=1

	//ROUTE: daftar barang (daftar package barang)
	router.HandleFunc("/barang", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			barang.GetBarang(w, r)
		} else if r.Method == http.MethodPost {
			barang.TambahBarang(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}).Methods("GET", "POST")

	// CORS agar Vue bisa akses backend
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	// http.HandleFunc("/barang", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == "GET" {
	// 		barang.GetBarang(w, r)
	// 	} else if r.Method == "POST" {
	// 		barang.TambahBarang(w, r)
	// 	}
	// })

	log.Println("Server running on :3000")
	log.Fatal(http.ListenAndServe(":3000", c.Handler(router)))
}

func connectDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/inventaris_barang")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var user User
	err = db.QueryRow("SELECT id, nama, username, role, status, password FROM user WHERE username = ?", req.Username).
		Scan(&user.ID, &user.Nama, &user.Username, &user.Role, &user.Status, &user.Password)

	if err != nil || req.Password != user.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	user.Password = "" // jangan kirim password
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user":  user,
		"token": "dummy-token",
	})
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, nama, username, role, status FROM user")
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Nama, &u.Username, &u.Role, &u.Status)
		users = append(users, u)
	}

	json.NewEncoder(w).Encode(users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Gagal hapus user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	var u User
	json.NewDecoder(r.Body).Decode(&u)

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE user SET nama=?, username=?, role=?, status=? WHERE id=?",
		u.Nama, u.Username, u.Role, u.Status, id)

	if err != nil {
		http.Error(w, "Gagal update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Data tidak valid", http.StatusBadRequest)
		return
	}

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO user (nama, username, role, status, password) VALUES (?, ?, ?, ?, ?)",
		user.Nama, user.Username, user.Role, user.Status, user.Password)

	if err != nil {
		http.Error(w, "Gagal tambah user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
