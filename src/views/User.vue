<template>
  <div class="user-page">
    <h2 v-if="isAdmin">Manajemen User</h2>
    <h2 v-else>Profil Saya</h2>

    <!-- ADMIN VIEW -->
    <div v-if="isAdmin">
      <div class="header-bar">
        <input v-model="searchQuery" placeholder="Cari user..." class="search-input" />
        <button class="btn-tambah" @click="tambahUser">+ Tambah User</button>
      </div>

      <div class="table-container">
        <table class="user-table">
          <thead>
            <tr>
              <th>No</th>
              <th>Nama</th>
              <th>Username</th>
              <th>Role</th>
              <th>Status</th>
              <th>Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(user, index) in filteredUsers" :key="user.id">
              <td>{{ index + 1 }}</td>
              <td>{{ user.nama }}</td>
              <td>{{ user.username }}</td>
              <td>{{ user.role }}</td>
              <td>{{ user.status }}</td>
              <td>
                <button class="btn-edit" @click="editUser(user.id)">Edit</button>
                <button class="btn-delete" @click="hapusUser(user.id)">Hapus</button>
              </td>
            </tr>
            <tr v-if="filteredUsers.length === 0">
              <td colspan="6" class="text-center">Tidak ada user ditemukan</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- NON-ADMIN VIEW -->
    <div v-else class="user-profile">
      <div class="profile-box">
        <p><strong>Nama:</strong> {{ currentUser.nama }}</p>
        <p><strong>Username:</strong> {{ currentUser.username }}</p>
        <p><strong>Role:</strong> {{ currentUser.role }}</p>
        <p><strong>Status:</strong> {{ currentUser.status }}</p>
        <button class="btn-edit" @click="editProfil">Edit Profil</button>
        <button class="btn-reset" @click="resetPassword">Ubah Password</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'UserPage',
  data() {
    return {
      currentUser: {
        id: 99,
        nama: "Ketua Yayasan",
        username: "yayasan01",
        role: "ketua_yayasan",
        status: "Aktif"
      },
      searchQuery: '',
      users: [
        { id: 1, nama: "Admin", username: "admin1", role: "admin", status: "Aktif" },
        { id: 2, nama: "Ketua Masjid", username: "masjid01", role: "ketua_masjid", status: "Aktif" },
        { id: 3, nama: "Ketua Yayasan", username: "yayasan01", role: "ketua_yayasan", status: "Nonaktif" }
      ]
    };
  },
  computed: {
    isAdmin() {
      return this.currentUser.role === "admin";
    },
    filteredUsers() {
      const q = this.searchQuery.toLowerCase();
      return this.users.filter(user =>
        user.nama.toLowerCase().includes(q) ||
        user.username.toLowerCase().includes(q) ||
        user.role.toLowerCase().includes(q)
      );
    }
  },
  methods: {
    tambahUser() {
      alert("Navigasi ke form tambah user");
    },
    editUser(id) {
      alert(`Edit user ID ${id}`);
    },
    hapusUser(id) {
      if (confirm("Yakin ingin menghapus user ini?")) {
        this.users = this.users.filter(user => user.id !== id);
      }
    },
    editProfil() {
      alert("Edit profil saya");
    },
    resetPassword() {
      alert("Ubah password saya");
    }
  }
};
</script>

<style scoped>
.user-page {
  padding: 20px;
}

.header-bar {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
  gap: 10px;
}

.search-input {
  flex: 1;
  min-width: 200px;
  padding: 8px;
  border-radius: 4px;
  border: 1px solid #ccc;
}

.btn-tambah {
  background-color: #27ae60;
  color: white;
  padding: 10px 16px;
  border: none;
  border-radius: 4px;
  font-weight: bold;
  cursor: pointer;
}

.btn-tambah:hover {
  background-color: #219150;
}

.table-container {
  width: 100%;
  overflow-x: auto;
}

.user-table {
  width: 100%;
  border-collapse: collapse;
  min-width: 800px;
}

.user-table th,
.user-table td {
  padding: 10px 12px;
  border: 1px solid #ddd;
  text-align: left;
  vertical-align: middle;
}

.user-table th {
  background-color: #2c3e50;
  color: white;
}

.text-center {
  text-align: center;
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

.btn-reset {
  background-color: #f39c12;
  color: white;
  padding: 6px 10px;
  border: none;
  border-radius: 4px;
  margin-top: 10px;
  cursor: pointer;
}

.user-profile .profile-box {
  max-width: 500px;
  background-color: #f9f9f9;
  padding: 20px;
  border-radius: 8px;
  border: 1px solid #ccc;
}
</style>
