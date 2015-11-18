require 'vnehm/tracks_view_command'

module Vnehm

  ##
  # This command gets likes/posts from user's account,
  # Prints as menu, and downloads selected tracks

  class SelectCommand < TracksViewCommand

    def initialize
      super

      add_option(:to, 'to ПУТЬ',
                 'Скачивать аудиозапись(и) в ПУТЬ')

      add_option(:pl, 'pl ПЛЕЙЛИСТ',
                 "Добавлять аудиозапись(и) в плейлист iTunes'a c именем" \
                                                                  ' ПЛЕЙЛИСТ')

      add_option(:limit, 'limit ЧИСЛО',
                 'Показывать ЧИСЛО аудиозаписей на каждой страницe')

      add_option(:offset, 'offset ЧИСЛО',
                 'Показывать с ЧИСЛО+1 аудиозаписи')

      add_option(:dl, 'dl yes',
                 'Не добавлять аудиозапись(и) в iTunes. Просто скачать их')
    end

    def program_name
      'vnehm select'
    end

    def summary
      'Вывод ваших аудиозаписей из VK и загрузка выбранных треков'
    end

    def usage
      "#{program_name} [ОПЦИИ]"
    end

    protected

    def get_tracks
      @track_manager.tracks(@limit, @offset)
    end

  end
end
