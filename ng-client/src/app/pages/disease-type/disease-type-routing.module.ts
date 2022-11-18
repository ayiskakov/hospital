import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { DiseaseTypeComponent } from './disease-type.component';

const routes: Routes = [
  { path: '', component: DiseaseTypeComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class DiseaseTypeRoutingModule { }
