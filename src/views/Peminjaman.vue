<template>
  <div class="peminjaman">
    <h2>Daftar Peminjaman</h2>
    <div class="table-container">
        <table class="peminjaman-table">
          <thead>
            <tr>
              <th>No</th>
              <th>Kode Peminjam</th>
              <th>Nama Peminjam</th>
              <th>Nama Barang</th>
              <th>Merk</th>
              <th>Deskripsi</th>
              <th>Tgl Peminjaman</th>
              <th>Jatuh Tempo</th>
              <th>Tgl Pengembalian</th>
              <th>Status</th>
              <th>Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) in daftarPeminjaman" :key="item.id">
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

    <h2 style="margin-top: 40px;">Riwayat Peminjaman</h2>
    <div class="table-container">
        <table class="peminjaman-table">
          <thead>
            <tr>
              <th>No</th>
              <th>Kode Peminjam</th>
              <th>Nama Peminjam</th>
              <th>Nama Barang</th>
              <th>Deskripsi</th>
              <th>Tgl Peminjaman</th>
              <th>Jatuh Tempo</th>
              <th>Tgl Pengembalian</th>
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
  </div>
</template>

<script>
export default {
  name: 'Peminjaman',
  data() {
    return {
      daftarPeminjaman: [
        {
          id: 1,
          kode: 'PMJ001',
          nama: 'Ahmad',
          barang: 'Proyektor',
          merk: 'Epson',
          deskripsi: 'Digunakan untuk presentasi kelas',
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
          deskripsi: 'Diperlukan untuk kegiatan seminar',
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
          deskripsi: 'Kegiatan dokumentasi',
          tglPinjam: '2025-04-10',
          jatuhTempo: '2025-04-15',
          tglKembali: '2025-04-14',
          status: 'Dikembalikan'
        }
      ]
    }
  },
  methods: {
    editPeminjaman(id) {
      alert(`Edit peminjaman ID: ${id}`);
    },
    hapusPeminjaman(id) {
      if (confirm('Yakin ingin menghapus data peminjaman ini?')) {
        this.daftarPeminjaman = this.daftarPeminjaman.filter(item => item.id !== id);
      }
    }
  }
}
</script>

<style scoped>
.table-container {
    width: 100%;
    overflow-x: auto;
    margin-top: 16px;
    border-radius: 6px;
    background: white;
    box-shadow: 0 2px 5x rgba(0, 0, 0, 0.05);
}

.peminjaman-table {
    min-width: 1000px;
    width: 100%;
    border-collapse: collapse;
}

.peminjaman-table th,
.peminjaman-table td {
    padding: 12px;
    text-align: left;
    border-bottom: 1px solid #ddd;
}

.peminjaman-table th {
    background-color: #2c3e50;
    color: white;
    position: sticky;
    top: 0;
}

.peminjaman {
  padding: 20px;
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
}

.peminjaman h2 {
  margin: 20px 0 10px;
}

.peminjaman-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 10px;
}

.peminjaman-table th,
.peminjaman-table td {
  padding: 10px;
  border-bottom: 1px solid #ccc;
  text-align: left;
  vertical-align: top;
}

.peminjaman-table th {
  background-color: #2c3e50;
  color: white;
}

.btn-edit {
  background-color: #3498db;
  color: white;
  border: none;
  padding: 6px 10px;
  border-radius: 4px;
  margin-right: 5px;
  cursor: pointer;
}

.btn-delete {
  background-color: #e74c3c;
  color: white;
  border: none;
  padding: 6px 10px;
  border-radius: 4px;
  cursor: pointer;
}

.btn-edit:hover {
  background-color: #2980b9;
}

.btn-delete:hover {
  background-color: #c0392b;
}

@media (max-width: 768px) {
    .peminjaman-table {
        font-size: 14px;
    }
}
</style>

