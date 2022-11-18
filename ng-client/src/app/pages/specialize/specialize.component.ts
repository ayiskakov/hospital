import { Component, OnInit } from '@angular/core';
import { DiseaseTypeDto, DoctorDto, PublicServantDto, SpecializeDto } from 'src/app/core/app-api/dto';
import { AppApiService } from "../../core/app-api/app-api.service";
import { FormBuilder, FormGroup, Validators } from "@angular/forms";

@Component({
  selector: 'app-specialize',
  templateUrl: './specialize.component.html',
  styleUrls: ['./specialize.component.scss']
})
export class SpecializeComponent implements OnInit {
  specializeList: SpecializeDto[] = [];
  doctorList: DoctorDto[] = [];
  diseaseTypeList: DiseaseTypeDto[] = [];
  publicServantList: PublicServantDto[] = [];
  isVisible: boolean = false;
  loading: boolean = false;
  form: FormGroup;

  constructor(private api: AppApiService, private fb: FormBuilder) {
    this.form = this.fb.group({
      id: [null, [Validators.required]],
      email: [null, [Validators.required, Validators.email]],
    })
  }

  ngOnInit() {
    this.fetch();
    this.api.getDoctorList().subscribe({
      next: (res: any) => {
        if (res.doctors) {
          this.doctorList = res.doctors;
        } else {
          this.doctorList = [];
        }
      }
    })
    this.api.getDiseaseTypeList().subscribe({
      next: (res: any) => {
        if (res.disease_types) {
          this.diseaseTypeList = res.disease_types;
        } else {
          this.diseaseTypeList = [];
        }
      }
    })
  }

  showModal(): void {
    this.isVisible = true;
  }

  handleCancel(): void {
    this.isVisible = false;
  }

  fetch(): void {
    this.api.getSpecializeList().subscribe({
      next: (res: any) => {
        if (res.specializes) {
          this.specializeList = res.specializes;
        } else {
          this.specializeList = [];
        }
      }
    })
  }

  delete(specialize: SpecializeDto): void {
    this.api.deleteSpecialize(specialize.doctor.user.email, specialize.disease_type.id).subscribe({
      next: (res: any) => {
        this.fetch();
      }
    })
  }

  submitForm(form: SpecializeDto): void {
    console.log(form, "form")
    if (this.form.valid) {
      this.loading = true;
      this.api.createSpecialize(form).subscribe({
        next: (res: any) => {
          this.fetch();
          this.isVisible = false;
          this.loading = false;
        },
        error: () => {
          this.loading = false;
        }
      })
    } else {
      Object.values(this.form.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }
}
