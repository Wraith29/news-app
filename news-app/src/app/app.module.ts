import { NgModule } from "@angular/core";
import { AppComponent } from './app.component';
import { ArticleComponent } from "./article/article.component";
import { BrowserModule } from "@angular/platform-browser";
import { ArticleService } from "./article.service";
import { AppRoutingModule, routes } from "./app-routing.module";
import { FormsModule } from "@angular/forms";
import { AdminComponent } from './admin/admin.component';
import { HomeComponent } from './home/home.component';
import { provideRouter, RouterOutlet, withComponentInputBinding } from "@angular/router";
import { provideHttpClient } from "@angular/common/http";
import { provideAnimationsAsync } from '@angular/platform-browser/animations/async';
import { HeaderComponent } from './header/header.component';
import { MultiSelectModule } from "primeng/multiselect";
import { InputTextModule } from "primeng/inputtext";

@NgModule({
  declarations: [
    AdminComponent,
    AppComponent,
    ArticleComponent,
    HeaderComponent,
    HomeComponent,
  ],
  imports: [
    AppRoutingModule,
    BrowserModule,
    FormsModule,
    InputTextModule,
    MultiSelectModule,
    RouterOutlet,
  ],
  providers: [
    provideAnimationsAsync(),
    provideHttpClient(),
    provideRouter(routes, withComponentInputBinding()),
    ArticleService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
