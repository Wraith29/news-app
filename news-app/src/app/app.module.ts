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

@NgModule({
  declarations: [
    AppComponent,
    ArticleComponent,
    AdminComponent,
    HomeComponent,
    HeaderComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    RouterOutlet,
  ],
  providers: [
    provideHttpClient(),
    ArticleService,
    provideAnimationsAsync(),
    provideRouter(routes, withComponentInputBinding())
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
