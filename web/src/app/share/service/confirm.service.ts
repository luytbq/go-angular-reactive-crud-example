import {
  AfterViewInit,
  Component,
  ElementRef,
  inject,
  Injectable,
  ViewChild,
} from '@angular/core';
import {
  MAT_DIALOG_DATA,
  MatDialog,
  MatDialogConfig,
  MatDialogModule,
  MatDialogRef,
} from '@angular/material/dialog';
import { Observable } from 'rxjs/internal/Observable';

export type ConfirmServiceConfig = {
  title: string | null;
  prompt: string | null;
  yes: string | null;
  no: string | null;
};

@Injectable({
  providedIn: 'root',
})
export class ConfirmService {
  constructor() {}

  dialog = inject(MatDialog);
  dialogRef: MatDialogRef<ConfirmComponent> | undefined;

  result$: Observable<any> = new Observable();

  data: any = {};

  title(title: string): ConfirmService {
    this.data.title = title;
    return this;
  }

  prompt(prompt: string): ConfirmService {
    this.data.prompt = prompt;
    return this;
  }

  yes(yes: string): ConfirmService {
    this.data.yes = yes;
    return this;
  }

  no(no: string): ConfirmService {
    this.data.no = no;
    return this;
  }

  open() {
    this.dialogRef = this.dialog.open(ConfirmComponent, {
      data: this.data,
      hasBackdrop: true,
      width: '380px',
    });

    return this.dialogRef.afterClosed();
  }
}

@Component({
  selector: 'app-confirm',
  imports: [MatDialogModule],
  standalone: true,
  template: ` <div>
    <h2 class="p-3 text-center text-xl font-bold">
      {{ data.title || 'Confirm action' }}
    </h2>
    <hr/>
    <div class="mb-3 p-3" #dialog_body></div>
    <div class="flex flex-row-reverse gap-3 p-3">
      <button
        class="min-w-16 rounded border bg-green-500 p-2"
        mat-button
        mat-dialog-close
        (click)="closeDialog(true)"
      >
        {{ data.yes || 'Yes' }}
      </button>
      <button
        class="min-w-16 rounded border bg-gray-400 p-2"
        mat-button
        mat-dialog-close
        cdkFocusInitial
        (click)="closeDialog(false)"
      >
        {{ data.no || 'No' }}
      </button>
    </div>
  </div>`,
})
export class ConfirmComponent implements AfterViewInit {
  dialogRef = inject(MatDialogRef<ConfirmComponent>);
  data = inject(MAT_DIALOG_DATA);

  @ViewChild('dialog_body') dialog_body: ElementRef | undefined;

  constructor() {}

  ngAfterViewInit(): void {
    console.log(this.dialog_body);
    if (this.data.prompt) {
      this.dialog_body
        ? (this.dialog_body.nativeElement.innerHTML = this.data.prompt)
        : '';
    }
  }

  closeDialog(value: boolean) {
    this.dialogRef.close(value);
  }
}
