<template>
  <div class="peminjaman">

    <!-- === DAFTAR PEMINJAMAN === -->
     <div class="header-bar">
        <h2>Daftar Peminjaman Aktif</h2>
        <button class="btn-tambah" @click="tambahPeminjam">+ Tambah Peminjam</button>
     </div>

    <div class="search-bar">
      <input type="text" v-model="searchQuery" placeholder="Cari peminjam atau barang..." />
    </div>

    <!-- DESKTOP TABLE -->
    <div class="table-container desktop-only">
      <table class="peminjaman-table">
        <thead>
          <tr>
            <th>No</th>
            <th>Kode</th>
            <th>Nama</th>
            <th>Barang</th>
            <th>Merk</th>
            <th>Deskripsi</th>
            <th>Tgl Pinjam</th>
            <th>Jatuh Tempo</th>
            <th>Tgl Kembali</th>
            <th>Status</th>
            <th>Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in filteredPeminjaman" :key="item.id">
            <td>{{ index + 1 }}</td>
            <td>{{ item.kode }}</td>
            <td>{{ item.nama }}</td>
            <td>{{ item.barang }}</td>
            <td>{{ item.merk }}</td>
            <td>{{ item.deskripsi }}</td>
            <td>{{ item.tglPinjam }}</td>
            <td>{{ item.jatuhTempo }}</td>
            <td>{{ item.tglKembali || '-' }}</td>
            <td>{{ item.status }}</td>
            <td>
              <button @click="editPeminjaman(item.id)" class="btn-edit">Edit</button>
              <button @click="hapusPeminjaman(item.id)" class="btn-delete">Hapus</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- MOBILE CARD - Aktif -->
    <div class="mobile-card-container mobile-only">
      <div class="peminjaman-card" v-for="(item, index) in filteredPeminjaman" :key="item.id">
        <p><strong>No:</strong> {{ index + 1 }}</p>
        <p><strong>Kode:</strong> {{ item.kode }}</p>
        <p><strong>Nama:</strong> {{ item.nama }}</p>
        <p><strong>Barang:</strong> {{ item.barang }}</p>
        <p><strong>Merk:</strong> {{ item.merk }}</p>
        <p><strong>Deskripsi:</strong> {{ item.deskripsi }}</p>
        <p><strong>Tgl Pinjam:</strong> {{ item.tglPinjam }}</p>
        <p><strong>Jatuh Tempo:</strong> {{ item.jatuhTempo }}</p>
        <p><strong>Tgl Kembali:</strong> {{ item.tglKembali || '-' }}</p>
        <p><strong>Status:</strong> {{ item.status }}</p>
        <div class="card-actions">
          <button @click="editPeminjaman(item.id)" class="btn-edit">Edit</button>
          <button @click="hapusPeminjaman(item.id)" class="btn-delete">Hapus</button>
        </div>
      </div>
    </div>

    <!-- === RIWAYAT === -->
    <h2 class="section-title">Riwayat Peminjaman</h2>

    <!-- DESKTOP TABLE -->
    <div class="table-container desktop-only">
      <table class="peminjaman-table">
        <thead>
          <tr>
            <th>No</th>
            <th>Kode</th>
            <th>Nama</th>
            <th>Barang</th>
            <th>Deskripsi</th>
            <th>Tgl Pinjam</th>
            <th>Jatuh Tempo</th>
            <th>Tgl Kembali</th>
            <th>Status</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(riwayat, index) in riwayatPeminjaman" :key="riwayat.id">
            <td>{{ index + 1 }}</td>
            <td>{{ riwayat.kode }}</td>
            <td>{{ riwayat.nama }}</td>
            <td>{{ riwayat.barang }}</td>
            <td>{{ riwayat.deskripsi }}</td>
            <td>{{ riwayat.tglPinjam }}</td>
            <td>{{ riwayat.jatuhTempo }}</td>
            <td>{{ riwayat.tglKembali }}</td>
            <td>{{ riwayat.status }}</td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- MOBILE CARD - Riwayat -->
    <div class="mobile-card-container mobile-only riwayat-scroll">
      <div class="peminjaman-card" v-for="(riwayat, index) in riwayatPeminjaman" :key="riwayat.id">
        <p><strong>No:</strong> {{ index + 1 }}</p>
        <p><strong>Kode:</strong> {{ riwayat.kode }}</p>
        <p><strong>Nama:</strong> {{ riwayat.nama }}</p>
        <p><strong>Barang:</strong> {{ riwayat.barang }}</p>
        <p><strong>Deskripsi:</strong> {{ riwayat.deskripsi }}</p>
        <p><strong>Tgl Pinjam:</strong> {{ riwayat.tglPinjam }}</p>
        <p><strong>Jatuh Tempo:</strong> {{ riwayat.jatuhTempo }}</p>
        <p><strong>Tgl Kembali:</strong> {{ riwayat.tglKembali }}</p>
        <p><strong>Status:</strong> {{ riwayat.status }}</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      searchQuery: '',
      daftarPeminjaman: [
        {
          id: 1,
          kode: 'PMJ001',
          nama: 'Ahmad',
          barang: 'Proyektor',
          merk: 'Epson',
          deskripsi: 'Untuk presentasi',
          tglPinjam: '2025-05-20',
          jatuhTempo: '2025-05-25',
          tglKembali: '',
          status: 'Dipinjam'
        },
        {
          id: 2,
          kode: 'PMJ002',
          nama: 'Rina',
          barang: 'Laptop',
          merk: 'Asus',
          deskripsi: 'Untuk seminar',
          tglPinjam: '2025-05-22',
          jatuhTempo: '2025-05-27',
          tglKembali: '',
          status: 'Dipinjam'
        }
      ],
      riwayatPeminjaman: [
        {
          id: 101,
          kode: 'PMJ000',
          nama: 'Fajar',
          barang: 'Kamera',
          deskripsi: 'Dokumentasi',
          tglPinjam: '2025-04-10',
          jatuhTempo: '2025-04-15',
          tglKembali: '2025-04-14',
          status: 'Dikembalikan'
        },
        {
          id: 101,
          kode: 'PMJ000',
          nama: 'Fajar',
          barang: 'Kamera',
          deskripsi: 'Dokumentasi',
          tglPinjam: '2025-04-10',
          jatuhTempo: '2025-04-15',
          tglKembali: '2025-04-14',
          status: 'Dikembalikan'
        }
      ]
    };
  },
  computed: {
    filteredPeminjaman() {
      const q = this.searchQuery.toLowerCase();
      return this.daftarPeminjaman.filter(item =>
        item.nama.toLowerCase().includes(q) ||
        item.kode.toLowerCase().includes(q) ||
        item.barang.toLowerCase().includes(q)
      );
    }
  },
  methods: {
    tambahPeminjam() {
        alert('Navigasi ke form tambah peminjam');
    },
    editPeminjaman(id) {
      alert(`Edit peminjaman ID: ${id}`);
    },
    hapusPeminjaman(id) {
      if (confirm('Yakin ingin menghapus data ini?')) {
        this.daftarPeminjaman = this.daftarPeminjaman.filter(item => item.id !== id);
      }
    }
  }
};
</script>
<style scoped>
.peminjaman {
  max-height: 80vh;
  overflow-y: auto;  
  padding: 20px;
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
}

.header-bar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 10px;
}

.btn-tambah {
    background-color: #27ae60;
    color: white;
    border: none;
    padding: 10px 16px;
    border-radius: 4px;
    cursor: pointer;
    font-weight: bold;
    transition: background-color 0.3s ease;
}

.btn-tambah:hover {
    background-color: #219150;
}

/* .section-title {
  font-size: 1.2rem;
  margin: 30px 0 15px;
  padding-bottom: 5px;
  border-bottom: 2px solid #eee;
} */

.search-bar {
  margin-bottom: 16px;
}
.search-bar input {
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 6px;
}

.table-container {
  overflow-x: auto;
  margin-bottom: 20px;
}
.peminjaman-table {
  width: 100%;
  border-collapse: collapse;
}
.peminjaman-table th, .peminjaman-table td {
  padding: 10px;
  border-bottom: 1px solid #ddd;
  text-align: left;
}
.peminjaman-table th {
  background-color: #2c3e50;
  color: white;
}

.btn-edit, .btn-delete {
  padding: 6px 10px;
  border-radius: 4px;
  border: none;
  color: white;
  cursor: pointer;
}
.btn-edit {
  background-color: #3498db;
}
.btn-edit:hover {
  background-color: #2980b9;
}
.btn-delete {
  background-color: #e74c3c;
}
.btn-delete:hover {
  background-color: #c0392b;
}

.mobile-only {
  display: none;
}
.desktop-only {
  display: block;
}

@media (max-width: 900px) {
  .desktop-only {
    display: none;
  }
  .mobile-only {
    display: flex;
    flex-direction: column;
  }
  .mobile-card-container {
    max-height: 400px;
    overflow-y: auto;
    margin-bottom: 30px;
  }
  .peminjaman-card {
    background: #f8f9fa;
    border-radius: 8px;
    padding: 14px;
    margin-bottom: 12px;
    box-shadow: 0 2px 5px rgba(0,0,0,0.07);
  }
}
</style>
