import { Component, OnInit } from '@angular/core';
import { DiseaseTypeDto } from 'src/app/core/app-api/dto';
import { AppApiService } from "../../core/app-api/app-api.service";
import { FormBuilder, FormGroup, Validators } from "@angular/forms";

@Component({
  selector: 'app-disease-type',
  templateUrl: './disease-type.component.html',
  styleUrls: ['./disease-type.component.scss']
})
export class DiseaseTypeComponent implements OnInit {
  diseaseTypeList: DiseaseTypeDto[] = [];
  isVisible: boolean = false;
  isEditing: boolean = false;
  loading: boolean = false;
  form: FormGroup;

  constructor(private api: AppApiService, private fb: FormBuilder) {
    this.form = this.fb.group({
      description: [null, [Validators.required]]
    })
  }

  ngOnInit() {
    this.fetch();
  }

  showModal(): void {
    this.form = this.fb.group({
      description: [null, [Validators.required]]
    })
    this.isEditing = false;
    this.isVisible = true;
  }

  handleCancel(): void {
    this.isVisible = false;
  }

  view(diseaseType: DiseaseTypeDto) {
    this.form = this.fb.group({
      id: [diseaseType.id, [Validators.required]],
      description: [diseaseType.description, [Validators.required]]
    })
    this.isEditing = true;
    this.isVisible = true;
  }

  fetch(): void {
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

  delete(diseaseType: DiseaseTypeDto): void {
    this.api.deleteDiseaseType(diseaseType.id).subscribe({
      next: (res: any) => {
        this.fetch();
      }
    })
  }

  submitForm(form: DiseaseTypeDto): void {
    if (this.form.valid) {
      this.loading = true;
      const payload: DiseaseTypeDto = { id: +form.id, description: form.description };
      if (this.isEditing) {
        this.api.updateDiseaseType(payload.id, payload).subscribe({
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
        this.api.createDiseaseType(payload).subscribe({
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
