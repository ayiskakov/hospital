import { Component, OnInit } from '@angular/core';
import { CountryDto, PublicServantDto } from 'src/app/core/app-api/dto';
import { AppApiService } from "../../core/app-api/app-api.service";
import { FormBuilder, FormGroup, Validators } from "@angular/forms";

@Component({
  selector: 'app-public-servant',
  templateUrl: './public-servant.component.html',
  styleUrls: ['./public-servant.component.scss']
})
export class PublicServantComponent implements OnInit {
  publicServantList: PublicServantDto[] = [];
  countryList: CountryDto[] = [];
  isVisible: boolean = false;
  isEditing: boolean = false;
  loading: boolean = false;
  form: FormGroup;

  constructor(private api: AppApiService, private fb: FormBuilder) {
    this.form = this.fb.group({
      email: [null, [Validators.required, Validators.email]],
      department: [null, [Validators.required]],
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
  }

  showModal(): void {
    this.form = this.fb.group({
      email: [null, [Validators.required, Validators.email]],
      department: [null, [Validators.required]],
    })
    this.isEditing = false;
    this.isVisible = true;
  }

  handleCancel(): void {
    this.isVisible = false;
  }

  view(publicServant: any) {
    this.form = this.fb.group({
      email: [publicServant.user.email, [Validators.required, Validators.email]],
      department: [publicServant.department, [Validators.required]],
    })
    this.isEditing = true;
    this.isVisible = true;
  }

  fetch(): void {
    this.api.getPublicServantList().subscribe({
      next: (res: any) => {
        if (res.public_servants) {
          this.publicServantList = res.public_servants;
        } else {
          this.publicServantList = [];
        }
      }
    })
  }

  delete(publicServant: PublicServantDto): void {
    this.api.deletePublicServant(publicServant.user.email).subscribe({
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
        public_servant: {
          department: form.department
        }
      };
      if (this.isEditing) {
        this.api.updatePublicServant(payload.user.email, payload.public_servant).subscribe({
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
        this.api.createPublicServant(payload).subscribe({
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
