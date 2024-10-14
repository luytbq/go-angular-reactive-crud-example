import { Component, inject } from '@angular/core';
import { CategoryServiceService } from '../share/service/category-service.service';
import { FormControl, FormGroup, FormsModule } from '@angular/forms';
import { AddCategory, Category } from '../model/category';
import {
  MAT_DIALOG_DATA,
  MatDialog,
  MatDialogModule,
  MatDialogRef,
} from '@angular/material/dialog';
import { ConfirmService } from '../share/service/confirm.service';
import { Subject } from 'rxjs/internal/Subject';

@Component({
  selector: 'app-category',
  standalone: true,
  imports: [FormsModule, MatDialogModule],
  templateUrl: './category.component.html',
  styleUrl: './category.component.css',
})
export class CategoryComponent {
  service = inject(CategoryServiceService);
  confirm = inject(ConfirmService);
  dialog = inject(MatDialog);

  searchKeyword = '';
  addCategory: AddCategory = {
    name: '',
  };

  formAdd = new FormGroup({
    name: new FormControl(''),
  });

  dialogRefEdit: MatDialogRef<EditCategory> | undefined;

  constructor() {}

  search() {
    this.service.search$.next({ keyword: this.searchKeyword });
  }

  add() {
    this.dialog.open(EditCategory, {
      data: {
        mode: 'add',
      },
    }).afterClosed().subscribe((result) => {
      if (result) {
        this.service.add$.next(result);
      }
    });
  }

  delete(id: number, name: string) {
    this.confirm
      .title('Delete category?')
      .prompt(`Deleting <b>${name}</b><br/>This can\' be undone!`)
      .yes('Confirm')
      .no('Cancel')
      .open()
      .subscribe((result) => {
        if (result) {
          this.service.delete$.next(id);
        }
      });
  }

  edit(category: Category) {
    this.dialog.open(EditCategory, {
      data: {
        category,
        mode: 'edit',
      },
    }).afterClosed().subscribe((result) => {
      if (result) {
        this.service.edit$.next(result);
      }
    });
  }
}

@Component({
  standalone: true,
  imports: [FormsModule],
  template: ` <div>
    <h2 class="p-3 text-center text-xl font-bold">
      {{ (mode == 'edit' ? 'Edit ' : 'Add new') + ' category' }}
    </h2>
    <hr />
    <div class="mb-3 p-3">
      <div class="flex gap-3 items-center">
        <label class="inline-block">Name</label>
        <input
          class="inline-block rounded border px-3 py-2"
          [(ngModel)]="category.name"
        />
      </div>
    </div>
    <div class="flex flex-row-reverse gap-3 p-3">
      <button
        class="min-w-16 rounded border bg-green-500 p-2"
        mat-button
        mat-dialog-close
        (click)="save()"
      >Save</button>
      <button
        class="min-w-16 rounded border bg-gray-400 p-2"
        mat-button
        mat-dialog-close
        cdkFocusInitial
        (click)="cancel()"
      >Cancel</button>
    </div>
  </div>`,
})
class EditCategory {
  category: Partial<Category> = {};
  mode: 'edit' | 'add' = 'edit';

  dialogRef = inject(MatDialogRef<EditCategory>);
  data = inject(MAT_DIALOG_DATA);

  constructor() {
    // make a copy of the passed object
    this.category = { ...this.data.category };
    this.mode = this.data.mode;
  }

  cancel() {
    this.dialogRef.close(undefined);
  }

  save() {
    this.dialogRef.close(this.category);
  }
}
