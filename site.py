from flask import Flask, url_for, render_template, redirect
from flask.ext.assets import Environment, Bundle

app = Flask(__name__)
assets = Environment(app)

smint = Bundle('js/jquery.smint.js')
comments = Bundle('js/comments.js')

assets.register('smint', smint)
assets.register('comments', comments)

@app.route('/')
def reroute_to_home():
    return redirect(url_for('resume'))

@app.route('/home')
def home():
    return render_template('home.html')

@app.route('/code')
def code():
    return render_template('code.html')

@app.route('/resume')
def resume():
    return render_template('resume.html')

@app.route('/blog')
def blog():
    return render_template('blog.html')

@app.route('/contact')
def contact():
    return render_template('contact.html')

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=80)
