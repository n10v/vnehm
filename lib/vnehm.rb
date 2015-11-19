require 'colored'

require 'vnehm/cfg'
require 'vnehm/client'
require 'vnehm/command_manager'
require 'vnehm/os'
require 'vnehm/path_manager'
require 'vnehm/playlist_manager'
require 'vnehm/track_manager'
require 'vnehm/ui'

module Vnehm

  class VnehmExit < SystemExit; end

  def self.start(args)

    begin
      init unless initialized?

      if args.empty?
        UI.say HELP
        UI.term
      end

      CommandManager.run(args)
    rescue StandardError, Timeout::Error => ex
      Vnehm::UI.term "Ошибка во время исполнения ... (#{ex.class})\n    #{ex}"
    rescue Interrupt
    rescue VnehmExit
    end
  end

  HELP = <<-EOF
#{'vnehm'.green} - это консольная утилита, которая скачивает (и добавляет в Вашу библиотеку iTunes) аудиозаписи из ВКонтакте

#{'Доступные команды:'.yellow}
  #{'get'.green}        Загрузка и добавление треков из VK в Вашу библиотеку iTunes
  #{'dl'.green}         Загрузка Ваших аудиозаписей из VK
  #{'configure'.green}  Настройка приложения
  #{'help'.green}       Показ справки для определенной команды
  #{'list'.green}     Вывод списка Ваших аудиозаписей из VK и загрузка выбранных треков
  #{'search'.green}     Поиск, вывод и загрузка определенных аудиозаписей по запросу
  #{'version'.green}    Вывод версии приложения

Используйте #{'vnehm help КОМАНДА'.yellow}, чтобы узнать подробнее об определенной команде

Команды и аргументы (но НЕ опции) могут быть сокращены, насколько они могут быть однозначны
Например, #{'nehm g'.magenta} может быть сокращением для #{'nehm get'.magenta}
EOF

  module_function

  def init
    UI.say 'Прежде чем использовать vnehm, Вам нужно его настроить:'
    Cfg.create unless Cfg.exist?

    PathManager.set_dl_path
    UI.newline

    if OS.mac?
      PlaylistManager.set_playlist
      UI.newline
    end

    UI.say 'Теперь Вам нужно авторизоваться'
    UI.say 'Для этого следуйте инструкциям ниже:'
    TokenManager.authorize
    UI.newline

    UI.success "Теперь вы можете использовать vnehm!"
    UI.newline
    UI.sleep
  end

  def initialized?
    Cfg.exist?
  end

end
