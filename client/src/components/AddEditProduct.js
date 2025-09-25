// import React, { useState, useEffect } from 'react';
// import '../styles/AddEditProduct.css';

// function AddEditProduct({ product, onSave, onCancel }) {
//   const [formData, setFormData] = useState({
//     title: '',
//     description: '',
//     price: '',
//     imgUrl: '',
//   });

//   useEffect(() => {
//     if (product) {
//       setFormData({
//         title: product.title,
//         description: product.description,
//         price: product.price,
//         imgUrl: product.imgUrl,
//       });
//     }
//   }, [product]);

//   const handleChange = (e) => {
//     const { name, value } = e.target;
//     setFormData((prevData) => ({
//       ...prevData,
//       [name]: value,
//     }));
//   };

//   const handleSubmit = (e) => {
//     e.preventDefault();
//     onSave(formData);
//   };

//   return (
//     <div className="add-edit-product">
//       <h2>{product ? 'Edit Product' : 'Add Product'}</h2>
//       <form onSubmit={handleSubmit}>
//         <label>
//           Title:
//           <input
//             type="text"
//             name="title"
//             value={formData.title}
//             onChange={handleChange}
//           />
//         </label>
//         <label>
//           Description:
//           <textarea
//             name="description"
//             value={formData.description}
//             onChange={handleChange}
//           />
//         </label>
//         <label>
//           Price:
//           <input
//             type="number"
//             name="price"
//             value={formData.price}
//             onChange={handleChange}
//           />
//         </label>
//         <label>
//           Image URL:
//           <input
//             type="text"
//             name="imgUrl"
//             value={formData.imgUrl}
//             onChange={handleChange}
//           />
//         </label>
//         <div className="buttons">
//           <button type="submit">{product ? 'Save Changes' : 'Add Product'}</button>
//           <button type="button" onClick={onCancel}>
//             Cancel
//           </button>
//         </div>
//       </form>
//     </div>
//   );
// }

// export default AddEditProduct;

import React, { useState, useEffect } from 'react';
import '../styles/AddEditProduct.css';

function AddEditProduct({ product, onSave, onCancel }) {
  const [formData, setFormData] = useState({
    title: '',
    description: '',
    price: '',
    imgUrl: '',
  });

  useEffect(() => {
    if (product) {
      setFormData({
        title: product.title,
        description: product.description,
        price: product.price,
        imgUrl: product.imgUrl,
      });
    }
  }, [product]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
      [name]: name === 'price' ? parseFloat(value) || 0 : value, // ✅ convert
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    await onSave(formData);
    setFormData({
      title: '',
      description: '',
      price: '',
      imgUrl: '',
    });
  };

  return (
    <div className='add-edit-product'>
      <h2>{product ? 'Edit Product' : 'Add Product'}</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Title:
          <input
            type='text'
            name='title'
            value={formData.title}
            onChange={handleChange}
          />
        </label>
        <label>
          Description:
          <textarea
            name='description'
            value={formData.description}
            onChange={handleChange}
          />
        </label>
        <label>
          Price:
          <input
            type='number'
            name='price'
            value={formData.price}
            onChange={handleChange}
          />
        </label>
        <label>
          Image URL:
          <input
            type='text'
            name='imgUrl'
            value={formData.imgUrl}
            onChange={handleChange}
          />
        </label>
        <div className='buttons'>
          <button type='submit'>
            {product ? 'Save Changes' : 'Add Product'}
          </button>
          <button type='button' onClick={onCancel}>
            Cancel
          </button>
        </div>
      </form>
    </div>
  );
}

export default AddEditProduct;
