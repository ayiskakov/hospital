import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { SpecializeComponent } from './specialize.component';

const routes: Routes = [
  { path: '', component: SpecializeComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class SpecializeRoutingModule { }
