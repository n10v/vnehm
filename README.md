# vnehm

*vnehm* - это консольная утилита, которая скачивает (и добавляет в Вашу библиотеку iTunes) аудиозаписи из ВКонтакте

[![Gem Version](https://img.shields.io/gem/v/vnehm.svg)](https://rubygems.org/gems/vnehm)
[![Code Climate](https://img.shields.io/codeclimate/github/bogem/vnehm.svg)](https://codeclimate.com/github/bogem/vnehm)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/bogem/vnehm/blob/master/LICENSE)

## ОТКАЗ ОТ ОТВЕТСТВЕННОСТИ

***Используйте эту программу только для личного пользования***

***Разработчик vnehm не несет ответственности за любое нелегальное использование данной программы***

## Установка

**1. [Установите Ruby](https://www.ruby-lang.org/ru/downloads/)**

**2. Установите библиотеку `taglib`**

**Для Mac OS X:**

`brew install taglib`

or

`sudo port install taglib`

**Для Linux:**

Debian/Ubuntu: `sudo apt-get install libtag1-dev`

Fedora/RHEL: `sudo yum install taglib-devel`

**3. Установите `vnehm`**

`gem install vnehm`

## Перед использованием

Если Вы только что установили `vnehm`, введите любую команду для первоначальной инициализации

Например, `vnehm help`

`vnehm` должен ответить примерно так:
```
Прежде чем использовать vnehm, Вам нужно его настроить
Введите путь в желаемую директорию скачиваемых аудио...
```
А дальше следуйте инструкциям, которые вам предложит `vnehm`

## Примеры использования

Используйте `vnehm help` для списка всех доступных команд или `vnehm help КОМАНДА` для определенной команды

Команды и аргументы (но **НЕ** опции) могут быть сокращены, насколько они могут быть однозначны

#### Скачать в директорию по умолчанию и добавить в iTunes Вашу последнюю аудиозапись

  `$ vnehm get` = `$ vnehm g`

#### Скачать и добавить в iTunes несколько последних аудиозаписей

  `$ vnehm get 3` = `$ vnehm g 3`

#### Просто скачать аудиозапись

  `$ vnehm dl` = `$ vnehm d`

#### Скачать аудиозапись в другую директорию

  `$ vnehm g to ~/Downloads` or `$ vnehm d to .`

#### Скачать и добавить трек в другой плейлист iTunes

  `$ vnehm g pl MyPlaylist`

#### Вывести список Ваших аудиозаписей и скачать выбранные

  `$ vnehm list` = `$ vnehm l`

#### Найти треки по запросу и скачать выбранные

  `$ vnehm search kanye west` = `$ vnehm s kanye west`

## Лицензия

MIT
