import { Component, OnInit } from '@angular/core';
import { CountryDto } from 'src/app/core/app-api/dto';
import { AppApiService } from "../../core/app-api/app-api.service";
import { FormBuilder, FormGroup, Validators } from "@angular/forms";

@Component({
  selector: 'app-country',
  templateUrl: './country.component.html',
  styleUrls: ['./country.component.scss']
})
export class CountryComponent implements OnInit {
  countryList: CountryDto[] = [];
  isVisible: boolean = false;
  isEditing: boolean = false;
  loading: boolean = false;
  form: FormGroup;

  constructor(private api: AppApiService, private fb: FormBuilder) {
    this.form = this.fb.group({
      cname: [null, [Validators.required]],
      population: [null, [Validators.required, Validators.min(0)]]
    })
  }

  ngOnInit() {
    this.fetch();
  }

  showModal(): void {
    this.form = this.fb.group({
      cname: [null, [Validators.required]],
      population: [null, [Validators.required, Validators.min(0)]]
    })
    this.isEditing = false;
    this.isVisible = true;
  }

  handleCancel(): void {
    this.isVisible = false;
  }

  view(country: CountryDto) {
    this.form = this.fb.group({
      cname: [country.cname, [Validators.required]],
      population: [country.population, [Validators.required, Validators.min(0)]]
    })
    this.isEditing = true;
    this.isVisible = true;
  }

  fetch(): void {
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

  delete(country: CountryDto): void {
    this.api.deleteCountry(country.cname).subscribe({
      next: (res: any) => {
        this.fetch();
      }
    })
  }

  submitForm(form: CountryDto): void {
    if (this.form.valid) {
      this.loading = true;
      const payload: CountryDto = { cname: form.cname, population: +form.population };
      if (this.isEditing) {
        this.api.updateCountry(payload.cname, payload).subscribe({
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
        this.api.createCountry(payload).subscribe({
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
