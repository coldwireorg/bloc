<template>
  <div class="fx-column explorer-container" style="gap: 10px">
    <p>Explorer</p>
    <table>
      <thead>
        <tr>
          <th>
            Name
          </th>
          <th>Size</th>
          <th>Last modification</th>
          <th>Shared with</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td class="fx-row" style="align-items: center; gap: 8px;">
            <Icon
              name="folder-explorer24x24"
              size="24"
            />
            Wikileaks Case
          </td>
          <td>Maria Anders</td>
          <td>Germany</td>
          <td />
        </tr>
        <tr>
          <td class="fx-row" style="align-items: center; gap: 8px;">
            <Icon
              name="binder-white32x32"
              bg-color="var(--secondary-green)"
              size="10"
              padding="5"
              margin="2"
              brd-radius="5"
            />
            Centro comercial Moctezuma
          </td>
          <td>Francisco Chang</td>
          <td>Mexico</td>
          <td />
        </tr>
        <tr>
          <td class="fx-row" style="align-items: center; gap: 8px;">
            <Icon
              name="binder-white32x32"
              bg-color="var(--secondary-green)"
              size="10"
              padding="5"
              margin="2"
              brd-radius="5"
            />
            Centro comercial Moctezuma
          </td>
          <td>Francisco Chang</td>
          <td>Mexico</td>
          <td />
        </tr>
        <tr>
          <td class="fx-row" style="align-items: center; gap: 8px;">
            <Icon
              name="binder-white32x32"
              bg-color="var(--secondary-green)"
              size="10"
              padding="5"
              margin="2"
              brd-radius="5"
            />
            Centro comercial Moctezuma
          </td>
          <td>Francisco Chang</td>
          <td>Mexico</td>
          <td />
        </tr>
        <tr>
          <td class="fx-row" style="align-items: center; gap: 8px;">
            <Icon
              name="binder-white32x32"
              bg-color="var(--secondary-green)"
              size="10"
              padding="5"
              margin="2"
              brd-radius="5"
            />
            Centro comercial Moctezuma
          </td>
          <td>Francisco Chang</td>
          <td>Mexico</td>
          <td />
        </tr>
        <tr>
          <td class="fx-row" style="align-items: center; gap: 8px;">
            <Icon
              name="binder-white32x32"
              bg-color="var(--secondary-green)"
              size="10"
              padding="5"
              margin="2"
              brd-radius="5"
            />
            Centro comercial Moctezuma
          </td>
          <td>Francisco Chang</td>
          <td>Mexico</td>
          <td />
        </tr>
        <tr>
          <td class="fx-row" style="align-items: center; gap: 8px;">
            <Icon
              name="binder-white32x32"
              bg-color="var(--secondary-green)"
              size="10"
              padding="5"
              margin="2"
              brd-radius="5"
            />
            Centro comercial Moctezuma
          </td>
          <td>Francisco Chang</td>
          <td>Mexico</td>
          <td />
        </tr>
        <tr>
          <td class="fx-row" style="align-items: center; gap: 8px;">
            <Icon
              name="binder-white32x32"
              bg-color="var(--secondary-green)"
              size="10"
              padding="5"
              margin="2"
              brd-radius="5"
            />
            Centro comercial Moctezuma
          </td>
          <td>Francisco Chang</td>
          <td>Mexico</td>
          <td />
        </tr>
        <tr>
          <td class="fx-row" style="align-items: center; gap: 8px;">
            <Icon
              name="binder-white32x32"
              bg-color="var(--secondary-green)"
              size="10"
              padding="5"
              margin="2"
              brd-radius="5"
            />
            Centro comercial Moctezuma
          </td>
          <td>Francisco Chang</td>
          <td>Mexico</td>
          <td />
        </tr>
        <tr>
          <td class="fx-row" style="align-items: center; gap: 8px;">
            <Icon
              name="binder-white32x32"
              bg-color="var(--secondary-green)"
              size="10"
              padding="5"
              margin="2"
              brd-radius="5"
            />
            Centro comercial Moctezuma
          </td>
          <td>Francisco Chang</td>
          <td>Mexico</td>
          <td />
        </tr>
        <tr>
          <td class="fx-row" style="align-items: center; gap: 8px;">
            <Icon
              name="binder-white32x32"
              bg-color="var(--secondary-green)"
              size="10"
              padding="5"
              margin="2"
              brd-radius="5"
            />
            Centro comercial Moctezuma
          </td>
          <td>Francisco Chang</td>
          <td>Mexico</td>
          <td />
        </tr>
        <tr>
          <td class="fx-row" style="align-items: center; gap: 8px;">
            <Icon
              name="binder-white32x32"
              bg-color="var(--secondary-green)"
              size="10"
              padding="5"
              margin="2"
              brd-radius="5"
            />
            Centro comercial Moctezuma
          </td>
          <td>Francisco Chang</td>
          <td>Mexico</td>
          <td />
        </tr>
        <tr>
          <td class="fx-row" style="align-items: center; gap: 8px;">
            <Icon
              name="binder-white32x32"
              bg-color="var(--secondary-green)"
              size="10"
              padding="5"
              margin="2"
              brd-radius="5"
            />
            Centro comercial Moctezuma
          </td>
          <td>Francisco Chang</td>
          <td>Mexico</td>
          <td />
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
  .explorer-container {
    height: 100%;
  }

  .explorer-container table {
    width: 100%;
    box-sizing: border-box;
    border-collapse: collapse;
    border-spacing: 0;
    text-align: left;
  }

  tbody {
    overflow: auto;
    height: 100%;
    width: 100%;
  }

  thead {
    border-bottom: 1px solid var(--complementary-white-50);
    color: var(--complementary-white-50);
    font-size: 10px;
    height: 32px;
  }

  tbody:before {
    content:"@";
    display:block;
    line-height:10px;
    text-indent:-99999px;
  }

  table tbody tr:hover {
    background-color: var(--complementary-white-5);
  }

  table td {
    font-size: 12px;
    height: 54px;
  }
</style>
