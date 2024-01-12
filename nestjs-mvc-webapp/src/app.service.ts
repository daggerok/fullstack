import { Injectable } from '@nestjs/common';
import { HandlebarsService } from '@gboutte/nestjs-hbs';

@Injectable()
export class AppService {
  constructor(private readonly hbs: HandlebarsService) {}
  renderIndexPage(): string {
    return this.hbs.renderFile('index.hbs', { name: 'Max' });
  }
}
