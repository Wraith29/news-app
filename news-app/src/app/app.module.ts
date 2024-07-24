import { NgModule } from "@angular/core";
import { AppComponent } from "./app.component";
import { ArticleComponent } from "./article/article.component";
import { BrowserModule } from "@angular/platform-browser";
import { AppRoutingModule, routes } from "./app-routing.module";
import { FormsModule } from "@angular/forms";
import { AdminComponent } from "./admin/admin.component";
import { HomeComponent } from "./home/home.component";
import {
  provideRouter,
  RouterOutlet,
  withComponentInputBinding,
} from "@angular/router";
import { provideHttpClient, withFetch } from "@angular/common/http";
import { provideAnimationsAsync } from "@angular/platform-browser/animations/async";
import { HeaderComponent } from "./header/header.component";
import { MultiSelectModule } from "primeng/multiselect";
import { InputTextModule } from "primeng/inputtext";
import { CacheService } from "./services/cache.service";
import { ArticleService } from "./services/article.service";
import { AuthorService } from "./services/author.service";
import { FeedService } from "./services/feed.service";
import { TableModule } from "primeng/table";
import { ButtonModule } from "primeng/button";
import { ConfirmDialogModule } from "primeng/confirmdialog";
import { ConfirmationService } from "primeng/api";

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
    ButtonModule,
    ConfirmDialogModule,
    FormsModule,
    InputTextModule,
    MultiSelectModule,
    RouterOutlet,
    TableModule,
  ],
  providers: [
    provideAnimationsAsync(),
    provideHttpClient(withFetch()),
    provideRouter(routes, withComponentInputBinding()),
    ArticleService,
    AuthorService,
    CacheService,
    FeedService,
    ConfirmationService,
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
