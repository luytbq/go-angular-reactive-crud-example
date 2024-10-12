import { Component, inject } from '@angular/core';
import { CategoryServiceService } from '../share/category-service.service';
import { FormControl, FormGroup, FormsModule } from '@angular/forms';
import { AddCategory, Category } from '../model/category';

@Component({
  selector: 'app-category',
  standalone: true,
  imports: [
    FormsModule,
  ],
  templateUrl: './category.component.html',
  styleUrl: './category.component.css',
})
export class CategoryComponent {
  service = inject(CategoryServiceService);

  searchKeyword = ''
  addCategory: AddCategory = {
    name: ''
  }

  formAdd = new FormGroup({
    name: new FormControl(''),
  });

  submitSearch() {
    this.service.search$.next({keyword: this.searchKeyword});
  }

  submitAdd() {
    this.service.add$.next(this.addCategory);
    this.addCategory.name = '';
  }

  delete(id: number) {
    this.service.delete$.next(id);
  }
}
