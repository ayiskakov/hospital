import { NgModule } from '@angular/core';

import { DoctorRoutingModule } from './doctor-routing.module';

import { DoctorComponent } from './doctor.component';
import { NzTableModule } from "ng-zorro-antd/table";
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { ReactiveFormsModule } from "@angular/forms";
import { NzFormModule } from "ng-zorro-antd/form";
import { NzInputModule } from "ng-zorro-antd/input";
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzSelectModule } from 'ng-zorro-antd/select';
import { NzInputNumberModule } from 'ng-zorro-antd/input-number';
import { CommonModule } from "@angular/common";


@NgModule({
  imports: [
    DoctorRoutingModule,
    NzTableModule,
    NzButtonModule,
    NzModalModule,
    ReactiveFormsModule,
    NzFormModule,
    NzInputModule,
    NzInputNumberModule,
    NzIconModule,
    NzSelectModule,
    CommonModule,
  ],
  declarations: [DoctorComponent],
  exports: [DoctorComponent],
})
export class DoctorModule { }
