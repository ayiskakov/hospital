import { NgModule } from '@angular/core';

import { PublicServantRoutingModule } from './public-servant-routing.module';

import { PublicServantComponent } from './public-servant.component';
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
    PublicServantRoutingModule,
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
  declarations: [PublicServantComponent],
  exports: [PublicServantComponent],
})
export class PublicServantModule { }
