import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: '/countries' },
  {
    path: 'countries', loadChildren: () => import('./pages/country/country.module').then(m => m.CountryModule)
  },
  {
    path: 'users', loadChildren: () => import('./pages/user/user.module').then(m => m.UserModule)
  },
  {
    path: 'doctors', loadChildren: () => import('./pages/doctor/doctor.module').then(m => m.DoctorModule)
  },
  {
    path: 'public-servants', loadChildren: () => import('./pages/public-servant/public-servant.module').then(m => m.PublicServantModule)
  },
  {
    path: 'disease-types', loadChildren: () => import('./pages/disease-type/disease-type.module').then(m => m.DiseaseTypeModule)
  },
  {
    path: 'diseases', loadChildren: () => import('./pages/disease/disease.module').then(m => m.DiseaseModule)
  },
  {
    path: 'discovery', loadChildren: () => import('./pages/discovery/discovery.module').then(m => m.DiscoveryModule)
  },
  {
    path: 'records', loadChildren: () => import('./pages/record/record.module').then(m => m.RecordModule)
  },
  {
    path: 'specialize', loadChildren: () => import('./pages/specialize/specialize.module').then(m => m.SpecializeModule)
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
