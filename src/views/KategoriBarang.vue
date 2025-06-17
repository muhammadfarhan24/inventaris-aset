<template>
  <div class="kategori-barang">
    <div class="header-bar">
      <h2>Daftar Kategori Barang</h2>
      <button class="btn-tambah" @click="tambahKategori">+ Tambah Kategori</button>
    </div>

    <div class="search-bar">
      <input type="text" v-model="searchQuery" placeholder="Cari Kategori..." />
    </div>

    <!-- Form Tambah / Edit -->
    <div v-if="showForm" class="form-kategori">
      <h3>{{ formMode === 'tambah' ? 'Tambah' : 'Edit' }} Kategori</h3>
      <input type="text" v-model="formData.kode_kategori" placeholder="Kode Kategori" disabled />
      <input type="text" v-model="formData.nama_kategori" placeholder="Nama Kategori" />
      <button @click="submitForm" class="btn-simpan">Simpan</button>
      <button @click="batalForm" class="btn-batal">Batal</button>
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
        <tr v-for="(kategori, index) in filteredKategoriList" :key="kategori.id">
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
  data() {
    return {
      searchQuery: '',
      kategoriList: [],
      formMode: 'tambah',
      showForm: false,
      formData: {
        id: null,
        kode_kategori: '',
        nama_kategori: ''
      }
    };
  },
  computed: {
    filteredKategoriList() {
      const q = this.searchQuery.toLowerCase();
      return this.kategoriList.filter(k =>
        k.nama_kategori.toLowerCase().includes(q) || k.kode_kategori.toLowerCase().includes(q)
      );
    }
  },
  methods: {
    fetchKategori() {
      fetch('http://localhost:3000/kategori')
        .then(res => res.json())
        .then(data => {
          this.kategoriList = data;
        });
    },
    generateKodeBaru() {
      const prefix = 'KT';
      const angkaList = this.kategoriList
        .map(k => parseInt(k.kode_kategori?.replace(prefix, '')))
        .filter(n => !isNaN(n));

      const maxAngka = angkaList.length > 0 ? Math.max(...angkaList) : 0;
      const nextAngka = maxAngka + 1;

      return prefix + nextAngka.toString().padStart(3, '0');
    },
    tambahKategori() {
      this.formMode = 'tambah';
      this.formData = {
        id: null,
        kode_kategori: this.generateKodeBaru(),
        nama_kategori: ''
      };
      this.showForm = true;
    },
    editKategori(id) {
      const kat = this.kategoriList.find(k => k.id === id);
      if (kat) {
        this.formMode = 'edit';
        this.formData = { ...kat };
        this.showForm = true;
      }
    },
    deleteKategori(id) {
      if (!confirm('Yakin ingin menghapus kategori ini?')) return;
      fetch(`http://localhost:3000/kategori/${id}`, {
        method: 'DELETE'
      }).then(() => this.fetchKategori());
    },
    submitForm() {
  const method = this.formMode === 'tambah' ? 'POST' : 'PUT';
  const url = this.formMode === 'tambah'
    ? 'http://localhost:3000/kategori'
    : `http://localhost:3000/kategori/${this.formData.id}`;

  fetch(url, {
    method,
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(this.formData)
  })
    .then(res => res.json())
    .then(() => {
      this.fetchKategori();
      this.showForm = false;
    });
},
    batalForm() {
      this.showForm = false;
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

.header-bar {
  display: flex;
  justify-content: space-between;
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

.search-bar {
  margin-bottom: 20px;
}
.search-bar input {
  width: 100%;
  padding: 10px 12px;
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
}
.kategori-table th {
  background-color: #2c3e50;
  color: white;
}

.btn-edit {
  background-color: #3498db;
  color: white;
  padding: 6px 10px;
  margin-right: 5px;
  border: none;
  border-radius: 4px;
}
.btn-delete {
  background-color: #e74c3c;
  color: white;
  padding: 6px 10px;
  border: none;
  border-radius: 4px;
}

.form-kategori {
  background: #f5f5f5;
  padding: 15px;
  margin-top: 20px;
  border-radius: 8px;
}
.form-kategori input {
  display: block;
  width: 100%;
  margin-bottom: 10px;
  padding: 8px;
}
.btn-simpan {
  background-color: #2ecc71;
  color: white;
  padding: 8px 16px;
  margin-right: 8px;
  border: none;
  border-radius: 4px;
}
.btn-batal {
  background-color: #95a5a6;
  color: white;
  padding: 8px 16px;
  border: none;
  border-radius: 4px;
}
</style>
