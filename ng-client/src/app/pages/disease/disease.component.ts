import { Component, OnInit } from '@angular/core';
import { DiseaseDto, DiseaseTypeDto } from 'src/app/core/app-api/dto';
import { AppApiService } from "../../core/app-api/app-api.service";
import { FormBuilder, FormGroup, Validators } from "@angular/forms";

@Component({
  selector: 'app-disease',
  templateUrl: './disease.component.html',
  styleUrls: ['./disease.component.scss']
})
export class DiseaseComponent implements OnInit {
  diseaseList: DiseaseDto[] = [];
  diseaseTypeList: DiseaseTypeDto[] = [];
  isVisible: boolean = false;
  isEditing: boolean = false;
  loading: boolean = false;
  form: FormGroup;

  constructor(private api: AppApiService, private fb: FormBuilder) {
    this.form = this.fb.group({
      id: [null, [Validators.required]],
      description: [null, [Validators.required]],
      disease_code: [null, [Validators.required]],
      pathogen: [null, [Validators.required]]
    })
  }

  ngOnInit() {
    this.fetch();
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
    this.form = this.fb.group({
      id: [null, [Validators.required]],
      description: [null, [Validators.required]],
      disease_code: [null, [Validators.required]],
      pathogen: [null, [Validators.required]]
    });
    this.isEditing = false;
    this.isVisible = true;
  }

  handleCancel(): void {
    this.isVisible = false;
  }

  view(disease: DiseaseDto) {
    this.form = this.fb.group({
      id: [disease.disease_type.id, [Validators.required]],
      description: [disease.description, [Validators.required]],
      disease_code: [disease.disease_code, [Validators.required]],
      pathogen: [disease.pathogen, [Validators.required]]
    })
    this.isEditing = true;
    this.isVisible = true;
  }

  fetch(): void {
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

  delete(disease: DiseaseDto): void {
    this.api.deleteDisease(disease.disease_code).subscribe({
      next: (res: any) => {
        this.fetch();
      }
    })
  }

  submitForm(form: DiseaseDto): void {
    if (this.form.valid) {
      this.loading = true;
      const payload: any = {
        id: +form.id,
        description: form.description,
        disease_code: form.disease_code,
        pathogen: form.pathogen
      };
      if (this.isEditing) {
        this.api.updateDisease(payload.disease_code, payload).subscribe({
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
        this.api.createDisease(payload).subscribe({
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
