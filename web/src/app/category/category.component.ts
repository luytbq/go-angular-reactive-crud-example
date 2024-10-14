import { Component, inject } from '@angular/core';
import { CategoryServiceService } from '../share/service/category-service.service';
import { FormControl, FormGroup, FormsModule } from '@angular/forms';
import { AddCategory } from '../model/category';
import { MatDialogModule } from '@angular/material/dialog';
import { ConfirmService } from '../share/service/confirm.service';
import { Subject } from 'rxjs/internal/Subject';

@Component({
  selector: 'app-category',
  standalone: true,
  imports: [
    FormsModule,
    MatDialogModule,
  ],
  templateUrl: './category.component.html',
  styleUrl: './category.component.css',
})
export class CategoryComponent {
  service = inject(CategoryServiceService);
  confirm = inject(ConfirmService);

  searchKeyword = ''
  addCategory: AddCategory = {
    name: ''
  }

  formAdd = new FormGroup({
    name: new FormControl(''),
  });

  delete$ = new Subject<number>();

  constructor() {
    this.delete$.subscribe(number => {
      this.confirm.title('Delete category?')
        .prompt('This can\' be undo')
        .yes('Confirm')
        .no('Cancel')
        .open().subscribe(result => {
          if (result) {
            this.service.delete$.next(number);
          }
        })
    })
  }

  search() {
    this.service.search$.next({keyword: this.searchKeyword});
  }

  add() {
    this.service.add$.next(this.addCategory);
    this.addCategory.name = '';
  }

  delete(id: number) {
    this.delete$.next(id);
  }
}
