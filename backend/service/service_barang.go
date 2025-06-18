package service

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type ServiceBarang struct {
	ID             int    `json:"id"`
	BarangID       int    `json:"barang_id"`
	NamaBarang     string `json:"nama_barang"`
	Deskripsi      string `json:"deskripsi"`
	TanggalService string `json:"tanggal_service"`
	Status         string `json:"status"`
}

func connectDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/inventaris_barang")
}

func GetServiceBarang(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT sb.id, sb.barang_id, b.nama, sb.deskripsi, sb.tanggal_service, sb.status
		FROM service_barang sb
		JOIN barang b ON sb.barang_id = b.id
	`)
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var list []ServiceBarang
	for rows.Next() {
		var s ServiceBarang
		// Perbaikan disini ⬇️
		if err := rows.Scan(&s.ID, &s.BarangID, &s.NamaBarang, &s.Deskripsi, &s.TanggalService, &s.Status); err != nil {
			continue
		}
		list = append(list, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func TambahServiceBarang(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var s ServiceBarang
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	s.Status = "Dalam Service"

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Masukkan ke service_barang
	_, err = db.Exec(`INSERT INTO service_barang (barang_id, deskripsi, tanggal_service, status)
		VALUES (?, ?, ?, ?)`, s.BarangID, s.Deskripsi, s.TanggalService, s.Status)
	if err != nil {
		http.Error(w, "Insert error", http.StatusInternalServerError)
		return
	}

	// Update status barang
	_, err = db.Exec("UPDATE barang SET status = 'Diservis' WHERE id = ?", s.BarangID)
	if err != nil {
		http.Error(w, "Gagal update status barang", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Barang masuk service & status diperbarui"})
}

func HapusServiceBarang(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM service_barang WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Delete error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Service barang berhasil dihapus"})
}

func SelesaiServiceBarang(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Ambil barang_id dari service_barang
	var barangID int
	err = db.QueryRow("SELECT barang_id FROM service_barang WHERE id = ?", id).Scan(&barangID)
	if err != nil {
		http.Error(w, "Gagal ambil data barang", http.StatusInternalServerError)
		return
	}

	// Update status di service_barang
	_, err = db.Exec("UPDATE service_barang SET status = 'Selesai' WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Gagal update status service", http.StatusInternalServerError)
		return
	}

	// Update barang: ubah status & kosongkan deskripsi
	_, err = db.Exec("UPDATE barang SET status = 'Tersedia', deskripsi = '' WHERE id = ?", barangID)
	if err != nil {
		http.Error(w, "Gagal update status barang", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Barang sudah selesai diservis dan tersedia kembali"})
}
