<template>
  <div class="kategori-barang">
    <div class="header-bar">
        <h2>Daftar Kategori Barang</h2>
        <button class="btn-tambah" @click="tambahKategori">+ Tambah Kategori</button>
    </div>

    <!-- Kolom Search -->
    <div class="search-bar">
        <input type="text" v-model="searchQuery" placeholder="Cari Kategori..." />
    </div> 

    <table class="kategori-table">
      <thead>
        <tr>
          <th>No</th>
          <th>Kode Kategori</th>
          <th>Kategori</th>
          <th>Aksi</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(kategori, index) in kategoriList" :key="kategori.id">
          <td>{{ index + 1 }}</td>
          <td>{{ kategori.kode_kategori }}</td>
          <td>{{ kategori.nama_kategori }}</td>
          <td>
            <button @click="editKategori(kategori.id)" class="btn-edit">Edit</button>
            <button @click="deleteKategori(kategori.id)" class="btn-delete">Delete</button>
          </td>
        </tr>

        <tr v-if="filteredKategoriList.length === 0">
            <td colspan="4" style="text-align: center;">Data Tidak Ditemukan</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: 'KategoriBarang',
  data() {
    return {
      searchQuery: '',
      kategoriList: []
    };
  },
  computed: {
    filteredKategoriList() {
    const query = this.searchQuery.toLowerCase();
    return this.kategoriList.filter(kat =>
    (kat.nama_kategori || '').toLowerCase().includes(query) ||
    (kat.kode_kategori || '').toLowerCase().includes(query)
  );
}

  },
  methods: {
    fetchKategori() {
      fetch('http://localhost:3000/kategori')
        .then(res => res.json())
        .then(data => {
          this.kategoriList = data;
        })
        .catch(err => {
          console.error('Gagal ambil data kategori:', err);
        });
    },
    tambahKategori() {
      const kode = prompt('Masukkan kode kategori:');
      const nama = prompt('Masukkan nama kategori:');
      if (!kode || !nama) return;

      fetch('http://localhost:3000/kategori', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ kode_kategori: kode, nama_kategori: nama })
      })
        .then(res => res.json())
        .then(() => {
          this.fetchKategori(); // refresh data
        })
        .catch(err => {
          console.error('Gagal tambah kategori:', err);
        });
    },
    deleteKategori(id) {
      if (!confirm('Yakin ingin menghapus kategori ini?')) return;

      fetch(`http://localhost:3000/kategori?id=${id}`, {
        method: 'DELETE'
      })
        .then(() => {
          this.fetchKategori(); // refresh data
        })
        .catch(err => {
          console.error('Gagal hapus kategori:', err);
        });
    },
    editKategori(id) {
      alert(`Edit kategori dengan ID: ${id}`);
      // fitur ini bisa dikembangkan nanti
    }
  },
  mounted() {
    this.fetchKategori();
  }
};
</script>


<style scoped>
.kategori-barang {
  padding: 20px;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
}

.kategori-barang h2 {
  margin-bottom: 20px;
}

.header-bar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
}

.btn-tambah {
    background-color: #27ae60;
    color: white;
    border: none;
    padding: 10px 16px;
    border-radius: 4px;
    cursor: pointer;
    font-weight: bold;
}

.btn-tambah:hover {
    background-color: #219150;
}

.search-bar {
    margin-bottom: 20px;
}

.search-bar input {
    width: 100%;
    padding: 10px 12px;
    font-size: 1rem;
    border: 1px solid #ccc;
    border-radius: 4px;
}

.kategori-table {
  width: 100%;
  border-collapse: collapse;
}

.kategori-table th,
.kategori-table td {
  padding: 12px;
  border-bottom: 1px solid #ccc;
  text-align: left;
}

.kategori-table th {
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
</style>
