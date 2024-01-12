import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ServeStaticModule } from '@nestjs/serve-static';
import { join } from 'path';
import { HandlebarsModule } from '@gboutte/nestjs-hbs';

@Module({
  imports: [
    ServeStaticModule.forRoot({
      rootPath: join(__dirname, '..', 'static'), // => $rootDir/static
    }),
    HandlebarsModule.forRoot({
      templateDirectory: 'templates', // => $rootDir/templates
      templateOptions: {},
      compileOptions: {},
    }),
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
