import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { DiseaseComponent } from './disease.component';

const routes: Routes = [
  { path: '', component: DiseaseComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class DiseaseRoutingModule { }
