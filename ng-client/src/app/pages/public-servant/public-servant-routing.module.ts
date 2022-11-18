import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { PublicServantComponent } from './public-servant.component';

const routes: Routes = [
  { path: '', component: PublicServantComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PublicServantRoutingModule { }
