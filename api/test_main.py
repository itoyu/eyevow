from unittest import main, TestCase
from main import app

class MainTestCase(TestCase):
  def test_main(self):
    print(app)

if __name__ == '__main__':
    main()