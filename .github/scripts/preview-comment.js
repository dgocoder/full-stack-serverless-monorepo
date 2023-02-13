const stackOutputs = require('../../.sst/outputs.json');

const generateMarkdown = () => {
  const stacks = [];
  Object.entries(stackOutputs).map(([key, value]) => {
    const split = key.split('-');
    const stackName = split?.[split.length - 1].replace('Stack', '');
    stacks.push(generateRow(stackName, value));
  });
  stacks.push(generateRow('Web', { PreviewUrl: process.env.PREVIEW_URL }));
  const header = `**The latest updates on monorepo stack**. Updated at: ${new Date().toUTCString()} 🎉\n| Name | Stack Outputs |\n| :--- | :------------ |\n${stacks.join(
    '\n'
  )}`;
  return header;
};

const generateRow = (stackName, outputs) => {
  const propertyRows = Object.entries(outputs).map(
    ([key, value]) => `<tr><td>${key}</td><td>${value}</td></tr>`
  );
  const template = `&nbsp;<table>  <thead>  <tr>  <th>Key</th>  <th>Values</th>  </tr>  </thead>  <tbody>  ${propertyRows.join(
    ' '
  )}  </tbody>  </table>`;
  const row = `| **${stackName}** | ${propertyRows.length > 0 ? template : 'No outputs'} |`;
  return row;
};

// console.log(generateMarkdown());

module.exports = generateMarkdown;
