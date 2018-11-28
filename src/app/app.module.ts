import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { DatePipe } from '@angular/common';
import { FileUploadModule } from 'ng2-file-upload';
import {
  MatButtonModule,
  MatIconModule,
  MatSidenavModule,
  MatToolbarModule,
  MatListModule,
  MatTableModule,
  MatSortModule,
  MatPaginatorModule,
  MatDialogModule,
  MatCardModule,
  MatGridListModule,
  MatProgressBarModule,
} from '@angular/material';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { ListComponent } from './list/list.component';
import { PeopleComponent } from './people/people.component';
import { SingleComponent } from './photo/single/single.component';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    ListComponent,
    PeopleComponent,
    SingleComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    HttpClientModule,
    MatButtonModule,
    MatIconModule,
    MatSidenavModule,
    MatToolbarModule,
    MatListModule,
    MatTableModule,
    MatSortModule,
    MatPaginatorModule,
    MatDialogModule,
    MatCardModule,
    MatGridListModule,
    MatProgressBarModule,

    FileUploadModule,
    AppRoutingModule,
  ],
  providers: [DatePipe],
  bootstrap: [AppComponent],
  entryComponents: [SingleComponent],
})
export class AppModule { }
