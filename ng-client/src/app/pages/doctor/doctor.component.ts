import { Component, OnInit } from '@angular/core';
import {CountryDto, DiseaseDto, DoctorDto, UserDto} from 'src/app/core/app-api/dto';
import { AppApiService } from "../../core/app-api/app-api.service";
import { FormBuilder, FormGroup, Validators } from "@angular/forms";

@Component({
  selector: 'app-doctor',
  templateUrl: './doctor.component.html',
  styleUrls: ['./doctor.component.scss']
})
export class DoctorComponent implements OnInit {
  doctorList: DoctorDto[] = [];
  countryList: CountryDto[] = [];
  userList: UserDto[] = [];
  isVisible: boolean = false;
  isEditing: boolean = false;
  loading: boolean = false;
  form: FormGroup;

  constructor(private api: AppApiService, private fb: FormBuilder) {
    this.form = this.fb.group({
      email: [null, [Validators.required, Validators.email]],
      degree: [null, [Validators.required]],
    })
  }

  ngOnInit() {
    this.fetch();
    this.api.getCountryList().subscribe({
      next: (res: any) => {
        if (res.countries) {
          this.countryList = res.countries;
        } else {
          this.countryList = [];
        }
      }
    })
    this.api.getUserList().subscribe({
      next: (res: any) => {
        if (res.users) {
          this.userList = res.users;
        } else {
          this.userList = [];
        }
      }
    })
  }

  showModal(): void {
    this.form = this.fb.group({
      email: [null, [Validators.required, Validators.email]],
      degree: [null, [Validators.required]],
    })
    this.isEditing = false;
    this.isVisible = true;
  }

  handleCancel(): void {
    this.isVisible = false;
  }

  view(doctor: any) {
    console.log(doctor)
    this.form = this.fb.group({
      email: [doctor.user.email, [Validators.required, Validators.email]],
      degree: [doctor.degree, [Validators.required]],
    })
    this.isEditing = true;
    this.isVisible = true;
  }

  fetch(): void {
    this.api.getDoctorList().subscribe({
      next: (res: any) => {
        if (res.doctors) {
          this.doctorList = res.doctors;
        } else {
          this.doctorList = [];
        }
      }
    })
  }

  delete(doctor: DoctorDto): void {
    this.api.deleteDoctor(doctor.user.email).subscribe({
      next: (res: any) => {
        this.fetch();
      }
    })
  }

  submitForm(form: any): void {
    if (this.form.valid) {
      this.loading = true;
      const payload: any = {
        user: {
          email: form.email,
        },
        doctor: {
          degree: form.degree
        }
      };
      if (this.isEditing) {
        this.api.updateDoctor(payload.user.email, payload.doctor).subscribe({
          next: (res: any) => {
            this.fetch();
            this.isVisible = false;
            this.isEditing = false;
            this.loading = false;
          },
          error: () => {
            this.loading = false;
          }
        })
      } else {
        this.api.createDoctor(payload).subscribe({
          next: (res: any) => {
            this.fetch();
            this.isVisible = false;
            this.loading = false;
          },
          error: () => {
            this.loading = false;
          }
        })
      }
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
