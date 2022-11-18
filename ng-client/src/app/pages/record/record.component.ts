import { Component, OnInit } from '@angular/core';
import { CountryDto, DiseaseDto, PublicServantDto, RecordDto } from 'src/app/core/app-api/dto';
import { AppApiService } from "../../core/app-api/app-api.service";
import { FormBuilder, FormGroup, Validators } from "@angular/forms";

@Component({
  selector: 'app-record',
  templateUrl: './record.component.html',
  styleUrls: ['./record.component.scss']
})
export class RecordComponent implements OnInit {
  recordList: RecordDto[] = [];
  countryList: CountryDto[] = [];
  diseaseList: DiseaseDto[] = [];
  selectedRecord: any;
  publicServantList: PublicServantDto[] = [];
  isVisible: boolean = false;
  isEditing: boolean = false;
  loading: boolean = false;
  form: FormGroup;

  constructor(private api: AppApiService, private fb: FormBuilder) {
    this.form = this.fb.group({
      email: [null, [Validators.required, Validators.email]],
      cname: [null, [Validators.required]],
      disease_code: [null, [Validators.required]],
      total_deaths: [null, [Validators.required]],
      total_patients: [null, [Validators.required]]
    })
  }

  ngOnInit() {
    this.fetch();
    this.api.getPublicServantList().subscribe({
      next: (res: any) => {
        if (res.public_servants) {
          this.publicServantList = res.public_servants;
        } else {
          this.publicServantList = [];
        }
      }
    })
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
      email: [null, [Validators.required, Validators.email]],
      cname: [null, [Validators.required]],
      disease_code: [null, [Validators.required]],
      total_deaths: [null, [Validators.required]],
      total_patients: [null, [Validators.required]]
    });
    this.isEditing = false;
    this.isVisible = true;
  }

  handleCancel(): void {
    this.isVisible = false;
  }

  view(record: RecordDto) {
    this.selectedRecord = {
      email: record.public_servant.user.email,
      cname: record.country.cname,
      disease_code: record.disease.disease_code
    };
    this.form = this.fb.group({
      email: [record.public_servant.user.email, [Validators.required, Validators.email]],
      cname: [record.country.cname, [Validators.required]],
      disease_code: [record.disease.disease_code, [Validators.required]],
      total_deaths: [record.total_deaths, [Validators.required]],
      total_patients: [record.total_patients, [Validators.required]]
    })
    this.isEditing = true;
    this.isVisible = true;
  }

  fetch(): void {
    this.api.getRecordList().subscribe({
      next: (res: any) => {
        if (res.records) {
          this.recordList = res.records;
        } else {
          this.recordList = [];
        }
      }
    })
  }

  delete(record: RecordDto): void {
    this.api.deleteRecord(record.public_servant.user.email, record.country.cname, record.disease.disease_code).subscribe({
      next: (res: any) => {
        this.fetch();
      }
    })
  }

  submitForm(form: RecordDto): void {
    if (this.form.valid) {
      this.loading = true;
      const payload: any = {
        email: form.email,
        cname: form.cname,
        disease_code: form.disease_code,
        total_deaths: +form.total_deaths,
        total_patients: +form.total_patients
      }
      if (this.isEditing) {
        this.api.updateRecord(this.selectedRecord.email, this.selectedRecord.cname, this.selectedRecord.disease_code, { total_deaths: +form.total_deaths, total_patients: +form.total_patients }).subscribe({
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
        this.api.createRecord(payload).subscribe({
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
