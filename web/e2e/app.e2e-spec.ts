import { GoibhniuWebPage } from './app.po';

describe('goibhniu-web App', () => {
  let page: GoibhniuWebPage;

  beforeEach(() => {
    page = new GoibhniuWebPage();
  });

  it('should display welcome message', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('Welcome to app!!');
  });
});
