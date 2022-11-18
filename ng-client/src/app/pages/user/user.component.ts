import { Component, OnInit } from '@angular/core';
import { CountryDto, UserDto } from 'src/app/core/app-api/dto';
import { AppApiService } from "../../core/app-api/app-api.service";
import { FormBuilder, FormGroup, Validators } from "@angular/forms";

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.scss']
})
export class UserComponent implements OnInit {
  userList: UserDto[] = [];
  countryList: CountryDto[] = [];
  isVisible: boolean = false;
  isEditing: boolean = false;
  loading: boolean = false;
  form: FormGroup;

  constructor(private api: AppApiService, private fb: FormBuilder) {
    this.form = this.fb.group({
      name: [null, [Validators.required]],
      surname: [null, [Validators.required]],
      email: [null, [Validators.required, Validators.email]],
      phone: [null, [Validators.required]],
      salary: [null, [Validators.required]],
      cname: [null, [Validators.required]]
    })
  }

  ngOnInit() {
    this.fetch();
    this.api.getCountryList().subscribe({
      next: (res: any) => {
        if (res.countries) {
          this.countryList = res.countries;
        } else {
          this.userList = [];
        }
      }
    })
  }

  showModal(): void {
    this.form = this.fb.group({
      name: [null, [Validators.required]],
      surname: [null, [Validators.required]],
      email: [null, [Validators.required, Validators.email]],
      phone: [null, [Validators.required]],
      salary: [null, [Validators.required]],
      cname: [null, [Validators.required]]
    });
    this.isEditing = false;
    this.isVisible = true;
  }

  handleCancel(): void {
    this.isVisible = false;
  }

  view(user: UserDto) {
    this.form = this.fb.group({
      name: [user.name, [Validators.required]],
      surname: [user.surname, [Validators.required]],
      email: [user.email, [Validators.required, Validators.email]],
      phone: [user.phone, [Validators.required]],
      salary: [user.salary, [Validators.required]],
      cname: [user.country.cname, [Validators.required]]
    })
    this.isEditing = true;
    this.isVisible = true;
  }

  fetch(): void {
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

  delete(user: UserDto): void {
    this.api.deleteUser(user.email).subscribe({
      next: (res: any) => {
        this.fetch();
      }
    })
  }

  submitForm(form: any): void {
    if (this.form.valid) {
      this.loading = true;
      const payload: any = {
        name: form.name,
        surname: form.surname,
        email: form.email,
        phone: form.phone,
        salary: +form.salary,
        country: {
          cname: form.cname
        }
      };
      if (this.isEditing) {
        this.api.updateUser(payload.email, payload).subscribe({
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
        this.api.createUser({ user: payload }).subscribe({
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
