module Vnehm
  class GetCommand < Command

    FIRST_TRACK = [1, 0]

    def initialize
      super

      add_option(:to, 'to ПУТЬ',
                 'Скачать аудиозапись(и) в ПУТЬ')

      add_option(:pl, 'pl ПЛЕЙЛИСТ',
                 'Добавлять аудиозапись(и) в плейлист iTunes с именем ПЛЕЙЛИСТ')
    end

    def execute
      track_manager = TrackManager.new(@options)

      UI.say 'Получение информации об аудиозаписи(ях)'
      arg = @options[:args].pop
      tracks =
        case arg
          when /^\d$/ # If arg is number
            track_manager.tracks(arg, 0)
          when nil
            track_manager.tracks(*FIRST_TRACK)
          else
            UI.term "Введен некорректный аргумент"
        end

      UI.term 'У Вас ещё нет аудиозаписей' if tracks.nil?

      track_manager.process_tracks(tracks)
    end

    def arguments
        { 'ЧИСЛО' => '(Необязательно) Скачать последние ЧИСЛО ваших аудиозаписей' }
    end

    def program_name
      'vnehm get'
    end

    def summary
      'Загрузка и добавление треков из VK в Вашу библиотеку iTunes'
    end

    def usage
      "#{program_name} [ЧИСЛО] [ОПЦИИ]"
    end

  end
end
