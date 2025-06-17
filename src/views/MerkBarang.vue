<template>
  <div class="merk-barang">
    <div class="header-bar">
      <h2>Daftar Merk Barang</h2>
      <button class="btn-tambah" @click="tambahMerk">+ Tambah Merk</button>
    </div>

    <!-- Kolom Search -->
    <div class="search-bar">
      <input type="text" v-model="searchQuery" placeholder="Cari Merk..." />
    </div>

    <!-- Form Tambah / Edit -->
    <div v-if="showForm" class="form-merk">
      <h3>{{ formMode === 'tambah' ? 'Tambah' : 'Edit' }} Merk</h3>
      <input type="text" v-model="formData.kode_merk" placeholder="Kode Merk" disabled />
      <input type="text" v-model="formData.nama_merk" placeholder="Nama Merk" />
      <button @click="submitForm" class="btn-simpan">Simpan</button>
      <button @click="batalForm" class="btn-batal">Batal</button>
    </div>

    <table class="merk-table">
      <thead>
        <tr>
          <th>No</th>
          <th>Kode Merk</th>
          <th>Nama Merk</th>
          <th>Aksi</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(merk, index) in filteredMerkList" :key="merk.id">
          <td>{{ index + 1 }}</td>
          <td>{{ merk.kode_merk }}</td>
          <td>{{ merk.nama_merk }}</td>
          <td>
            <button @click="editMerk(merk.id)" class="btn-edit">Edit</button>
            <button @click="deleteMerk(merk.id)" class="btn-delete">Delete</button>
          </td>
        </tr>

        <tr v-if="filteredMerkList.length === 0">
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
      merkList: [],
      formMode: 'tambah',
      showForm: false,
      formData: {
        id: null,
        kode_merk: '',
        nama_merk: ''
      }
    };
  },
  computed: {
    filteredMerkList() {
      const q = this.searchQuery.toLowerCase();
      return this.merkList.filter(m =>
        m.nama_merk.toLowerCase().includes(q) || m.kode_merk.toLowerCase().includes(q)
      );
    }
  },
  methods: {
    fetchMerk() {
      fetch('http://localhost:3000/merk')
        .then(res => res.json())
        .then(data => {
          this.merkList = data;
        });
    },
    generateKodeBaru() {
      const prefix = 'MK';
      const angkaList = this.merkList
        .map(m => parseInt(m.kode_merk?.replace(prefix, '')))
        .filter(n => !isNaN(n));

      const maxAngka = angkaList.length > 0 ? Math.max(...angkaList) : 0;
      const nextAngka = maxAngka + 1;

      return prefix + nextAngka.toString().padStart(3, '0');
    },
    tambahMerk() {
      this.formMode = 'tambah';
      this.formData = {
        id: null,
        kode_merk: this.generateKodeBaru(),
        nama_merk: ''
      };
      this.showForm = true;
    },
    editMerk(id) {
      const merk = this.merkList.find(m => m.id === id);
      if (merk) {
        this.formMode = 'edit';
        this.formData = { ...merk };
        this.showForm = true;
      }
    },
    deleteMerk(id) {
      if (!confirm('Yakin ingin menghapus merk ini?')) return;

      fetch(`http://localhost:3000/merk/${id}`, {
        method: 'DELETE'
      }).then(() => this.fetchMerk());
    },
    submitForm() {
      const method = this.formMode === 'tambah' ? 'POST' : 'PUT';
      const url = this.formMode === 'tambah'
        ? 'http://localhost:3000/merk'
        : `http://localhost:3000/merk/${this.formData.id}`;

      fetch(url, {
        method,
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(this.formData)
      })
        .then(res => res.json())
        .then(() => {
          this.fetchMerk();
          this.showForm = false;
        });
    },
    batalForm() {
      this.showForm = false;
    }
  },
  mounted() {
    this.fetchMerk();
  }
};
</script>

<style scoped>
.merk-barang {
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

.merk-table {
  width: 100%;
  border-collapse: collapse;
}
.merk-table th,
.merk-table td {
  padding: 12px;
  border-bottom: 1px solid #ccc;
}
.merk-table th {
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

.form-merk {
  background: #f5f5f5;
  padding: 15px;
  margin-top: 20px;
  border-radius: 8px;
}
.form-merk input {
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
