import { HttpClient, HttpParams } from '@angular/common/http';
import { computed, Injectable, signal } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import {
  AddCategory,
  Category,
  CategorySearchParams,
  CategorySearchResponse,
} from '../../model/category';
import { Subject } from 'rxjs/internal/Subject';
import {
  catchError,
  concatMap,
  EMPTY,
  exhaustMap,
  merge,
  mergeMap,
  of,
  startWith,
  switchMap,
  tap,
} from 'rxjs';
import { environment } from '../../environment/environment';
import { ToastrService } from 'ngx-toastr';

interface State {
  loaded: boolean;
  queryParams: CategorySearchParams;
  items: Category[];
  error: string | null;
}

@Injectable({
  providedIn: 'root',
})
export class CategoryServiceService {
  private state = signal<State>({
    loaded: false,
    queryParams: { keyword: '' },
    items: [],
    error: null,
  });

  // selectors
  loaded$ = computed(() => this.state().loaded);
  items$ = computed(() => this.state().items);
  error$ = computed(() => this.state().error);

  // source
  add$ = new Subject<AddCategory>();
  delete$ = new Subject<number>();
  edit$ = new Subject<Category>();
  search$ = new Subject<CategorySearchParams>();

  private searched$ = this.search$.pipe(
    concatMap((queryParams) => {
      this.state.update((state) => ({ ...state, queryParams: queryParams }));
      return of(true);
    }),
  );

  private categoryAdded$ = this.add$.pipe(
    concatMap((category) =>
      this.http.post(environment.API_URL + '/categories', category).pipe(
        tap((response) => {
          console.log(response);
          this.toastr.success('New category added');
        }),
        catchError((error) => this.handleError(error)),
      ),
    ),
  );

  private categoryDeleted$ = this.delete$.pipe(
    mergeMap((id) =>
      this.http.delete(environment.API_URL + '/categories/' + id).pipe(
        tap((response) => {
          console.log(response);
          this.toastr.success('Category deleted successfully');
        }),
        catchError((error) => this.handleError(error)),
      ),
    ),
  );

  private categoryEdited$ = this.edit$.pipe(
    exhaustMap((category) =>
      this.http.patch(environment.API_URL + '/categories', category).pipe(
        tap((response) => {
          console.log(response);
          this.toastr.success('Saved');
        }),
        catchError((error) => this.handleError(error)),
      ),
    ),
  );

  constructor(
    private http: HttpClient,
    private toastr: ToastrService,
  ) {
    // do search when these event occured
    merge(
      this.categoryAdded$,
      this.categoryDeleted$,
      this.categoryEdited$,
      this.searched$,
    )
      .pipe(
        startWith(null),
        switchMap(() =>
          this.http
            .get(
              `${environment.API_URL}/categories?keyword=${this.state().queryParams.keyword}`,
              {
                observe: 'body',
              },
            )
            .pipe(catchError((error) => this.handleError(error))),
        ),
        takeUntilDestroyed(),
      )
      .subscribe((response) => {
        this.state.update((state) => ({
          ...state,
          loaded: true,
          items: (response as CategorySearchResponse).items,
        }));
      });
  }

  private handleError(error: any) {
    this.state.update((state) => ({ ...state, error: error }));
    this.toastr.error(error.error?.error || 'Something went wrong');
    return EMPTY;
  }
}
