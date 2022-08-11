from abc import ABC, abstractmethod
from io import StringIO

list_content = StringIO()

class TextProcessor():
    def __init__(self, list_strategy):
        self._list_strategy = list_strategy

    def set_list_strategy(self, list_strategy):
        self._list_strategy = list_strategy

    def append_list(self, items=[]):
        self._list_strategy.start(list_content)
        for item in items:
            self._list_strategy.addListItem(list_content, item)
        self._list_strategy.end(list_content)

    def reset(self):
        list_content.truncate(0)
        list_content.seek(0)

    def output(self):
        content = list_content.getvalue()
        print(content)


class ListStrategyInterface(ABC):
    @abstractmethod
    def start(self, string):
        pass

    @abstractmethod
    def end(self, string):
        pass

    @abstractmethod
    def addListItem(self, string, item):
        pass


class MarkdownListStrategy(ListStrategyInterface):
    def start(self, string):
        pass

    def end(self, string):
        pass

    def addListItem(self, string, item):
        string.write(' * {}\n'.format(item))


class HtmlListStrategy(ListStrategyInterface):
    def start(self, string):
        string.write('<ul>\n')

    def end(self, string):
        string.write('</ul>\n')

    def addListItem(self, string, item):
        string.write(' <li> {} </li>\n'.format(item))

def main():
    text_processor = TextProcessor(MarkdownListStrategy())
    text_processor.append_list(['foo', 'Yao', 'Ling'])
    text_processor.output()

    text_processor.reset()
    text_processor.set_list_strategy(HtmlListStrategy())
    text_processor.append_list(['100', '5K', '$999'])
    text_processor.output()


if __name__ == '__main__':
    main()
