import { NgModule } from '@angular/core';

import { UserRoutingModule } from './user-routing.module';

import { UserComponent } from './user.component';
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
    UserRoutingModule,
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
  declarations: [UserComponent],
  exports: [UserComponent],
})
export class UserModule { }
