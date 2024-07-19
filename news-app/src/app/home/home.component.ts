import { Component, OnDestroy, OnInit } from "@angular/core";
import { Article } from "../types/article";
import { Subscription } from "rxjs";
import { ArticleService } from "../article.service";

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent implements OnInit, OnDestroy {
  public visibleArticles: Article[] = [];

  private _articles: Article[] = [];
  private _subscriptions: Subscription[] = [];
  private _query: string = '';

  constructor(
    private _articleService: ArticleService,
  ) { }

  public ngOnInit(): void {
    this._subscriptions.push(
      this._articleService.getAllArticles().subscribe({
        next: (articles: Article[]) => {
          this._articles = articles ?? [];
          this._filterArticles(this._query);
        },
        error: (err: Error) => {
          console.error(err);
        }
      })
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

  private _filterArticles(query: string): void {
    this.visibleArticles = this._articles.filter((article: Article) => this._objectContains(article, query.toLowerCase()));
  }

  private _objectContains(item: object, searchValue: string): boolean {
    if (searchValue.length === 0) {
      return true;
    }

    for (const value of Object.values(item)) {
      switch (typeof value) {
        case 'string':
          return value.toLowerCase().includes(searchValue);
        case 'object':
          return this._objectContains(value, searchValue);
        default:
          return false;
      }
    }

    return false;
  }
}
