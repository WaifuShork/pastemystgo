import React from 'react';
import clsx from 'clsx';
import styles from './HomepageFeatures.module.css';

const FeatureList = [
  {
    title: 'PasteMyst API',
    Svg: require('../../static/img/pastemystav.svg').default,
    description: (
      <>
      PasteMystGo provides low-level bindings for the <a href="https://paste.myst.rs/api-docs/index">PasteMyst API</a>.
      Built to be fast, safe, and reliable -- with nearly 100% test coverage.
      </>
    ),
  },
  {
    title: 'Focus on Speed',
    Svg: require('../../static/img/go.svg').default,
    description: (
      <>
      Just because the API rate-limits, doesn't mean your library has to be slow, PasteMystGo is designed with speed in mind
      so you can focus on what matters.
      </>
    ),
  },
];

function Feature({Svg, title, description}) {
  return (
    <div className={clsx('col col--6')}>
      <div className="text--center">
       {/*} <Svg className={styles.featureSvg} alt={title} />*/}
      </div>
      <div className="text--center padding-horiz--md">
        <h3>{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
