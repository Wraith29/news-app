import { Component, OnDestroy, OnInit } from "@angular/core";
import { Article } from "../types/article";
import { Subscription } from "rxjs";
import { MultiSelectChangeEvent } from "primeng/multiselect";
import { ArticleService } from "../services/article.service";
import { AuthorService } from "../services/author.service";
import { AuthService } from "../services/auth.service";
import { Router } from "@angular/router";

@Component({
  selector: "app-home",
  templateUrl: "./home.component.html",
  styleUrl: "./home.component.css",
})
export class HomeComponent implements OnInit, OnDestroy {
  public visibleArticles: Article[] = [];

  public authors: string[] = [];
  public selectedAuthors: string[] = [];

  private _articles: Article[] = [];
  private _subscriptions: Subscription[] = [];
  private _query: string = "";

  constructor(
    private _articleService: ArticleService,
    private _authorService: AuthorService,
    private _authService: AuthService,
    private _router: Router,
  ) {}

  public ngOnInit(): void {
    if (!this._authService.isLoggedIn()) {
      // Make sure the localStorage is cleaned out before continuing
      this._authService.logout();
      this._router.navigate(["login"]);
    }

    this._subscriptions.push(
      this._articleService.getAll().subscribe({
        next: (articles: Article[]) => {
          this._articles = articles ?? [];
          this._filterArticles(this._query);
        },
        error: (err: Error) => {
          console.error(err);
        },
      }),
    );

    this._subscriptions.push(
      this._authorService.getAll().subscribe({
        next: (authors: string[]) => {
          this.authors = authors;
        },
      }),
    );
  }

  public ngOnDestroy(): void {
    this._subscriptions.forEach((sub: Subscription) => sub.unsubscribe());
  }

  public handleInput(event: KeyboardEvent): void {
    const query = (event.target! as HTMLInputElement).value!;
    this._query = query;
    this._filterArticles(query);
  }

  public onChange(event: MultiSelectChangeEvent): void {
    this.selectedAuthors = <string[]>event.value;

    this._filterArticles(this._query);
  }

  private _filterArticles(query: string): void {
    this.visibleArticles = this._articles.filter((article: Article) => {
      const queryFilter = this._objectContains(article, query.toLowerCase());
      let authorFilter = true;

      if (this.selectedAuthors && this.selectedAuthors.length > 0) {
        authorFilter = this.selectedAuthors.includes(article.author);
      }

      return queryFilter && authorFilter;
    });
  }

  private _objectContains(item: object, searchValue: string): boolean {
    if (searchValue.length === 0) {
      return true;
    }

    for (const value of Object.values(item)) {
      switch (typeof value) {
        case "string":
          return value.toLowerCase().includes(searchValue);
        case "object":
          return this._objectContains(value, searchValue);
        default:
          return false;
      }
    }

    return false;
  }
}
