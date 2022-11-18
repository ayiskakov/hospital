import { Component, OnInit } from '@angular/core';
import { CountryDto, DiscoveryDto, DiseaseDto } from 'src/app/core/app-api/dto';
import { AppApiService } from "../../core/app-api/app-api.service";
import { FormBuilder, FormGroup, Validators } from "@angular/forms";

@Component({
  selector: 'app-discovery',
  templateUrl: './discovery.component.html',
  styleUrls: ['./discovery.component.scss']
})
export class DiscoveryComponent implements OnInit {
  selectedDiscovery: any;
  discoveryList: DiscoveryDto[] = [];
  countryList: CountryDto[] = [];
  diseaseList: DiseaseDto[] = [];
  isVisible: boolean = false;
  isEditing: boolean = false;
  loading: boolean = false;
  form: FormGroup;

  constructor(private api: AppApiService, private fb: FormBuilder) {
    this.form = this.fb.group({
      cname: [null, [Validators.required]],
      disease_code: [null, [Validators.required]],
      first_enc_date: [null, [Validators.required]],
    });
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
    this.api.getDiseaseList().subscribe({
      next: (res: any) => {
        if (res.diseases) {
          this.diseaseList = res.diseases;
        } else {
          this.diseaseList = [];
        }
      }
    })
  }

  showModal(): void {
    this.form = this.fb.group({
      cname: [null, [Validators.required]],
      disease_code: [null, [Validators.required]],
      first_enc_date: [null, [Validators.required]],
    });
    this.isEditing = false;
    this.isVisible = true;
  }

  handleCancel(): void {
    this.isVisible = false;
  }

  view(discovery: DiscoveryDto) {
    this.selectedDiscovery = {
      cname: discovery.country.cname,
      disease_code: discovery.disease.disease_code,
      first_enc_date: discovery.first_enc_date
    }
    this.form = this.fb.group({
      cname: [discovery.country.cname, [Validators.required]],
      disease_code: [discovery.disease.disease_code, [Validators.required]],
      first_enc_date: [discovery.first_enc_date, [Validators.required]],
    })
    this.isEditing = true;
    this.isVisible = true;
  }

  fetch(): void {
    this.api.getDiscoveryList().subscribe({
      next: (res: any) => {
        if (res.discoveries) {
          this.discoveryList = res.discoveries;
        } else {
          this.discoveryList = [];
        }
      }
    })
  }

  delete(discovery: DiscoveryDto): void {
    this.api.deleteDiscovery(discovery.country.cname, discovery.disease.disease_code).subscribe({
      next: (res: any) => {
        this.fetch();
      }
    })
  }

  submitForm(form: DiscoveryDto): void {
    if (this.form.valid) {
      this.loading = true;
      if (this.isEditing) {
        this.api.updateDiscovery(this.selectedDiscovery.cname, this.selectedDiscovery.disease_code, form).subscribe({
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
        this.api.createDiscovery(form).subscribe({
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
