<template>
    <div class="dashboard">
        <!-- Sidebar -->
         <aside class="sidebar" :class="{ open: sidebarOpen }">
            <h2 @click="closeSidebar">Yayasan <br>Al-Istiqomah</h2>
            <ul>
                <li>
                    <router-link to="/dashboard">Dashboard</router-link>
                </li>
                <li>
                     <!-- Barang Menu -->
                     <a href="#" @click.prevent="toggleBarangMenu">
                        Barang <span style="float: right;">{{ barangMenuOpen ? '▲' : '▼' }}</span>
                     </a>
                     <ul v-show="barangMenuOpen" class="submenu">
                        <li><router-link to="/dashboard/daftar-barang">Daftar Barang</router-link></li>
                        <li><router-link to="/dashboard/kategori-barang">Kategori Barang</router-link></li>
                        <li><router-link to="/dashboard/lokasi-barang">Lokasi Barang</router-link></li>
                        <li><router-link to="/dashboard/merk-barang">Merk Barang</router-link></li>
                        <li><router-link to="/dashboard/status-barang">Status Barang</router-link></li>
                        <li><router-link to="/dashboard/peminjaman">Peminjaman</router-link></li>
                        <li><router-link to="/dashboard/service-barang">Service Barang</router-link></li>
                     </ul>
                </li>
                <li>
                     <!-- Ruangan Menu -->
                     <a href="#" @click.prevent="toggleRuanganMenu">
                        Ruangan <span style="float: right;">{{ ruanganMenuOpen ? '▲' : '▼' }}</span>
                     </a>
                     <ul v-show="ruanganMenuOpen" class="submenu">
                        <li><a href="#">Daftar Ruangan</a></li>
                        <li><a href="#">Gudang Imam A</a></li>
                        <li><a href="#">Gudang Imam B</a></li>
                        <li><a href="#">Gudang Tangga A</a></li>
                        <li><a href="#">Gudang Tangga B</a></li>
                        <li><a href="#">Sekretariat</a></li>
                        <li><a href="#">Ruangan A Lantai 2</a></li>
                     </ul>
                </li>

                <li><a href="#">Users</a></li>
                <button @click="logout" class="logout-button">Logout</button>
            </ul>
         </aside>

         <!-- Main Content -->
          <main class="main-content">
            <header class="header">
                <button class="toogle-btn d-md-none" @click="sidebarOpen = !sidebarOpen">
                    <i class="bi bi-list" style="font-size: 1.5rem; color: #333;"></i>
                </button>
                <h1>Yayasan Al-Istiqomah</h1>
            </header>
            <section class="content">
                <router-view />
            </section>
          </main>
    </div>
</template>

<script>
export default {
    name: "Layout",
    data() {
        return {
            sidebarOpen: false,
            barangMenuOpen: false,
            ruanganMenuOpen: false
        };
    },
    methods: {
        logout() {
            this.$router.push('/');
        },
        toggleBarangMenu() {
            this.barangMenuOpen = !this.barangMenuOpen;
        },
        toggleRuanganMenu() {
            this.ruanganMenuOpen = !this.ruanganMenuOpen;
        },
        closeSidebar() {
            if (window.innerWidth <=768) {
                this.sidebarOpen = false;
            }
        }
    }
}
</script>

<style>
.sidebar {
    width: 220px;
    background-color: #2c3e50;
    color: white;
    padding: 20px;
    height: 100vh;
    overflow-y: auto;
    transition: transform 0.3s ease;
}

.sidebar::-webkit-scrollbar {
    background-color: rgba(255, 255, 255, 0.3);
    border-radius: 3px;
}

.sidebar::-webkit-scrollbar-thumb {
    background-color: rgba(255, 255, 255, 0.3);
    border-radius: 3px;
}

.sidebar h2 {
    margin-bottom: 20px;
    font-size: 22px;
}

.sidebar ul {
    list-style: none;
    padding: 0;
}

.sidebar li {
    margin: 10px 0;
}

.sidebar a {
    color: white;
    text-decoration: none;
    display: block;
    padding: 5px 10px;
    border-radius: 4px;
}

.sidebar a:hover {
    background-color: #34495e;
}

.submenu {
    margin-left: 10px;
    padding-left: 10px;
    border-left: 1px solid rgba(255, 255, 255, 0.2);
}

.submenu li {
    margin: 5px 0;
}

.logout-button {
    margin-top: 20px;
    background-color: #e74c3c;
    border: none;
    color: white;
    padding:  8px 12px;
    border-radius: 4px;
    cursor: pointer;
    width: 100%;
}

.logout-button:hover {
    background-color: #c0392b;
}
</style>