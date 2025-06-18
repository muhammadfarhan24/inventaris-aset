package barang

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Barang struct {
	ID         int    `json:"id"`
	KodeBarang string `json: "kode_barang"`
	Nama       string `json:"nama"`
	Kategori   int    `json:"kategori_id"`
	Merk       int    `json:"merk_id"`
	Status     string `json:"status"`
	Ruangan    int    `json:"ruangan_id"`
	Deskripsi  string `json:"deskripsi"`
	Created_at string `json:"created_at"`
}

func connectDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/inventaris_barang")
}

func GetBarang(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, kode_barang, nama, kategori_id, merk_id, status, ruangan_id, deskripsi, created_at FROM barang")
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var list []Barang
	for rows.Next() {
		var b Barang
		err := rows.Scan(&b.ID, &b.KodeBarang, &b.Nama, &b.Kategori, &b.Merk, &b.Status, &b.Ruangan, &b.Deskripsi, &b.Created_at)
		if err != nil {
			continue
		}
		list = append(list, b)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func TambahBarang(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var b Barang
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	//  Tambahkan ini:
	kodeBarang, err := generateKodeBarang(db)
	if err != nil {
		http.Error(w, "Gagal generate kode barang", http.StatusInternalServerError)
		return
	}

	//  Gunakan kodeBarang hasil generate, bukan dari request
	_, err = db.Exec(`
		INSERT INTO barang (kode_barang, nama, kategori_id, merk_id, status, ruangan_id, deskripsi)
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		kodeBarang, b.Nama, b.Kategori, b.Merk, b.Status, b.Ruangan, b.Deskripsi)

	if err != nil {
		http.Error(w, "Insert error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message":     "Barang berhasil ditambahkan",
		"kode_barang": kodeBarang,
	})
}

func generateKodeBarang(db *sql.DB) (string, error) {
	var lastKode string
	err := db.QueryRow("SELECT kode_barang FROM barang ORDER BY id DESC LIMIT 1").Scan(&lastKode)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	prefix := "BRG"
	var nextNum int
	if lastKode != "" {
		_, err := fmt.Sscanf(lastKode, "BRG%04d", &nextNum)
		if err != nil {
			nextNum = 0
		}
	}
	nextNum++

	return fmt.Sprintf("%s%04d", prefix, nextNum), nil
}

func EditBarang(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var b Barang
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Kosongkan deskripsi jika status bukan "Rusak"
	if b.Status != "Rusak" {
		b.Deskripsi = ""
	}

	_, err = db.Exec(`
		UPDATE barang 
		SET nama = ?, kategori_id = ?, merk_id = ?, status = ?, ruangan_id = ?, deskripsi = ?
		WHERE id = ?`,
		b.Nama, b.Kategori, b.Merk, b.Status, b.Ruangan, b.Deskripsi, b.ID)

	if err != nil {
		http.Error(w, "Update error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Barang berhasil diperbarui",
	})
}
