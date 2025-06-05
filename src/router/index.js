import { createRouter, createWebHistory } from 'vue-router'
import LoginPage from '../views/Login.vue'
import Layout from '@/components/icons/Layout.vue'
import Dashboard from '@/views/Dashboard.vue'
import DaftarBarang from '@/views/DaftarBarang.vue'
import KategoriBarang from '@/views/KategoriBarang.vue'
import LokasiBarang from '@/views/LokasiBarang.vue'
import MerkBarang from '@/views/MerkBarang.vue'
import Peminjaman from '@/views/Peminjaman.vue'
import StatusBarang from '@/views/StatusBarang.vue'
import ServiceBarang from '@/views/ServiceBarang.vue'
import GudangImamA from '@/views/GudangImamA.vue'
import GudangImamB from '@/views/GudangImamB.vue'
import GudangTanggaA from '@/views/GudangTanggaA.vue'
import GudangTanggaB from '@/views/GudangTanggaB.vue'
import Sekretariat from '@/views/sekretariat.vue'
import Lantai2 from '@/views/Lantai2.vue'
import User from '@/views/User.vue'

const routes = [
  {
    path: '/',
    name: 'Login',
    component: LoginPage
  },
  {
    path: '/dashboard',
    component: Layout,
    children: [
      {
        path: '',
        name: 'Dashboard',
        component: Dashboard
      },
      {
        path: 'daftar-barang',
        name: 'DaftarBarang',
        component: DaftarBarang
      },
      {
        path: 'kategori-barang',
        name: 'KategoriBarang',
        component: KategoriBarang
      },
      {
        path: 'lokasi-barang',
        name: 'LokasiBarang', 
        component: LokasiBarang
      },
      {
        path: 'merk-barang',
        name: 'MerkBarang',
        component: MerkBarang
      },
      {
        path: 'peminjaman',
        name: 'Peminjaman',
        component: Peminjaman
      },
      {
        path: 'status-barang',
        name: 'StatusBarang',
        component: StatusBarang
      },
      {
        path: 'service-barang',
        name: 'Service-Barang',
        component: ServiceBarang
      },
      {
        path: 'gudang-imam-a',
        name: 'Gudang-Imam-A',
        component: GudangImamA
      },
      {
        path: 'gudang-imam-b',
        name: 'Gudang-Imam-B',
        component: GudangImamB
      },
      {
        path: 'gudang-tangga-a',
        name: 'Gudang-Tangga-A',
        component: GudangTanggaA
      },
      {
        path: 'gudang-tangga-b',
        name: 'Gudang-Tangga-B',
        component: GudangTanggaB
      },
      {
        path: 'sekretariat',
        name: 'Sekretariat',
        component: Sekretariat
      },
      {
        path: 'gudang-lantai-2',
        name: 'Gudang-Lantai-2',
        component: Lantai2
      },
      {
        path: 'user',
        name: 'User',
        component: User
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
